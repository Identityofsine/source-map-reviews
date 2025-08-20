-- +goose Up
-- Migration: Add review performance indexes
-- Description: Creates indexes to optimize review queries, filtering, and full-text search

-- Index for review content trigram search (for partial text matching)
CREATE INDEX IF NOT EXISTS idx_map_reviews_review_trigram 
ON map_reviews USING gin (review gin_trgm_ops) WHERE review IS NOT NULL;

-- Composite index for map reviews ordered by rating (stars) and date
CREATE INDEX IF NOT EXISTS idx_map_reviews_map_stars_created 
ON map_reviews (map_name, stars DESC, created_at DESC);

-- Index for finding reviews by specific star ratings
CREATE INDEX IF NOT EXISTS idx_map_reviews_stars_created 
ON map_reviews (stars, created_at DESC);

-- Composite index for user's reviews ordered by date
CREATE INDEX IF NOT EXISTS idx_map_reviews_reviewer_created 
ON map_reviews (reviewer, created_at DESC);

-- Index for finding maps with no reviews (used in search filters)
CREATE INDEX IF NOT EXISTS idx_map_reviews_map_name_covering 
ON map_reviews (map_name) INCLUDE (map_review_id, stars, created_at);

-- Index for average rating calculations per map
CREATE INDEX IF NOT EXISTS idx_map_reviews_map_name_stars 
ON map_reviews (map_name, stars);

-- Partial index for high-rated reviews (4-5 stars)
CREATE INDEX IF NOT EXISTS idx_map_reviews_high_rated 
ON map_reviews (map_name, created_at DESC) WHERE stars >= 4;

-- Partial index for low-rated reviews (1-2 stars)
CREATE INDEX IF NOT EXISTS idx_map_reviews_low_rated 
ON map_reviews (map_name, created_at DESC) WHERE stars <= 2;

-- Index for recent reviews (useful for trending/recent activity)
CREATE INDEX IF NOT EXISTS idx_map_reviews_recent 
ON map_reviews (created_at DESC, map_name);

-- +goose Down
-- Drop indexes in reverse order
DROP INDEX IF EXISTS idx_map_reviews_recent;
DROP INDEX IF EXISTS idx_map_reviews_low_rated;
DROP INDEX IF EXISTS idx_map_reviews_high_rated;
DROP INDEX IF EXISTS idx_map_reviews_map_name_stars;
DROP INDEX IF EXISTS idx_map_reviews_map_name_covering;
DROP INDEX IF EXISTS idx_map_reviews_reviewer_created;
DROP INDEX IF EXISTS idx_map_reviews_stars_created;
DROP INDEX IF EXISTS idx_map_reviews_map_stars_created;
DROP INDEX IF EXISTS idx_map_reviews_review_trigram;
