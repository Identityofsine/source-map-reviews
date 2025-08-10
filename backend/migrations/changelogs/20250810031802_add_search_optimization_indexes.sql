-- +goose Up
-- Migration: Add specialized search optimization indexes
-- Description: Creates advanced indexes for search analytics, performance monitoring, and specialized queries

-- Create a materialized view for map statistics (refreshed periodically)
-- This pre-calculates expensive aggregations for better search performance
CREATE MATERIALIZED VIEW IF NOT EXISTS map_search_stats AS
SELECT 
    m.map_name,
    m.created_at as map_created_at,
    COUNT(DISTINCT mr.map_review_id) as review_count,
    COALESCE(AVG(mr.stars::numeric), 0) as avg_rating,
    COUNT(DISTINCT mt.lk_tag) as tag_count,
    MAX(mr.created_at) as last_review_date,
    -- Search rank calculation for relevance scoring
    (COUNT(DISTINCT mr.map_review_id) * 0.4 + 
     COALESCE(AVG(mr.stars::numeric), 0) * 0.6) as search_rank
FROM maps m
LEFT JOIN map_reviews mr ON m.map_name = mr.map_name
LEFT JOIN map_tags mt ON m.map_name = mt.map_name
GROUP BY m.map_name, m.created_at;

-- Index on the materialized view for fast lookups
CREATE UNIQUE INDEX IF NOT EXISTS idx_map_search_stats_map_name 
ON map_search_stats (map_name);

-- Index for ranking-based searches (most popular/highest rated)
CREATE INDEX IF NOT EXISTS idx_map_search_stats_rank 
ON map_search_stats (search_rank DESC, review_count DESC);

-- Index for filtering by review count
CREATE INDEX IF NOT EXISTS idx_map_search_stats_review_count 
ON map_search_stats (review_count DESC, avg_rating DESC);

-- Index for tag-based popularity
CREATE INDEX IF NOT EXISTS idx_map_search_stats_tags 
ON map_search_stats (tag_count DESC, search_rank DESC);

-- Advanced text search combinations
-- Index for combined map name and review text search
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_combined_text_search 
ON map_reviews USING gin (
    (map_name || ' ' || COALESCE(review, '')) gin_trgm_ops
);

-- Index for search by map name with review quality filter
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_quality_search 
ON map_reviews (map_name, stars) 
WHERE stars >= 3 AND review IS NOT NULL AND length(review) > 10;

-- Index for finding maps with recent activity
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_recent_activity 
ON map_reviews (map_name, updated_at DESC) 
WHERE updated_at > created_at;

-- Index for user engagement tracking
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_user_engagement 
ON map_reviews (reviewer) 
INCLUDE (map_name, stars, created_at) 
WHERE created_at > (NOW() - INTERVAL '90 days');

-- Index for seasonal/temporal analysis
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_temporal 
ON map_reviews (EXTRACT(YEAR FROM created_at), EXTRACT(MONTH FROM created_at), map_name);

-- Index for finding controversial maps (high variance in ratings)
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_rating_analysis 
ON map_reviews (map_name, stars) 
INCLUDE (created_at, map_review_id);

-- Search performance optimization index
-- This index helps with pagination in search results
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_search_pagination 
ON maps (map_name) 
INCLUDE (created_at, updated_at);

-- Index for advanced tag filtering (for complex tag combinations)
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_advanced_tag_filtering 
ON map_tags (lk_tag) 
INCLUDE (map_name, created_at) 
WHERE lk_tag IN (
    SELECT lk_tag FROM map_tags 
    GROUP BY lk_tag 
    HAVING COUNT(*) >= 5  -- Only popular tags
);

-- Create function to refresh materialized view (can be called periodically)
CREATE OR REPLACE FUNCTION refresh_map_search_stats()
RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY map_search_stats;
END;
$$ LANGUAGE plpgsql;

-- +goose Down
-- Drop function and materialized view
DROP FUNCTION IF EXISTS refresh_map_search_stats();
DROP MATERIALIZED VIEW IF EXISTS map_search_stats;

-- Drop indexes in reverse order
DROP INDEX IF EXISTS idx_advanced_tag_filtering;
DROP INDEX IF EXISTS idx_maps_search_pagination;
DROP INDEX IF EXISTS idx_maps_rating_analysis;
DROP INDEX IF EXISTS idx_maps_temporal;
DROP INDEX IF EXISTS idx_user_engagement;
DROP INDEX IF EXISTS idx_maps_recent_activity;
DROP INDEX IF EXISTS idx_maps_quality_search;
DROP INDEX IF EXISTS idx_combined_text_search;
