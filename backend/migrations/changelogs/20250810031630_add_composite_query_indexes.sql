-- +goose Up
-- Migration: Add composite query performance indexes
-- Description: Creates composite indexes to optimize complex multi-table queries and common search patterns

-- Composite index for tag-based searches (optimizes the tag filtering subquery)
CREATE INDEX IF NOT EXISTS idx_map_tags_lk_tag_map_name 
ON map_tags (lk_tag, map_name);

-- Reverse composite index for map-based tag lookups
CREATE INDEX IF NOT EXISTS idx_map_tags_map_name_lk_tag 
ON map_tags (map_name, lk_tag);

-- Index for tag popularity analysis
CREATE INDEX IF NOT EXISTS idx_map_tags_lk_tag_created 
ON map_tags (lk_tag, created_at DESC);

-- Composite index for map review images (optimizes bulk image loading)
CREATE INDEX IF NOT EXISTS idx_map_review_images_review_image 
ON map_review_images (map_review_id, image_id);

-- Reverse index for finding reviews by image
CREATE INDEX IF NOT EXISTS idx_map_review_images_image_review 
ON map_review_images (image_id, map_review_id);

-- Index for image caption searches (if users search image descriptions)
CREATE INDEX IF NOT EXISTS idx_images_caption_trigram 
ON images USING gin (caption gin_trgm_ops) WHERE caption IS NOT NULL;

-- Composite index for maps with their creation info and path
CREATE INDEX IF NOT EXISTS idx_maps_created_path 
ON maps (created_at DESC, map_path, map_name);

-- Index to support the "reviewed/unreviewed" filter efficiently
-- This creates a unique list of reviewed map names for fast lookups
CREATE INDEX IF NOT EXISTS idx_map_reviews_distinct_map_names 
ON map_reviews (map_name) WHERE map_review_id IS NOT NULL;

-- Composite index for user activity tracking (reviews + timestamps)
CREATE INDEX IF NOT EXISTS idx_map_reviews_user_activity 
ON map_reviews (reviewer, map_name, created_at DESC);

-- Index for finding maps by tag count (popular tags)
-- This helps with queries that want maps with multiple tags
CREATE INDEX IF NOT EXISTS idx_map_tags_count_optimization 
ON map_tags (map_name) INCLUDE (lk_tag);

-- Index for tag-based map discovery (finding similar maps)
CREATE INDEX IF NOT EXISTS idx_map_tags_tag_discovery 
ON map_tags (lk_tag) INCLUDE (map_name, created_at);

-- Partial index for recently tagged maps (trending tags)
CREATE INDEX IF NOT EXISTS idx_map_tags_recent 
ON map_tags (lk_tag, created_at DESC);

-- Index for maps with high review activity
CREATE INDEX IF NOT EXISTS idx_maps_review_activity 
ON map_reviews (map_name) INCLUDE (created_at, stars);

-- +goose Down
-- Drop indexes in reverse order
DROP INDEX IF EXISTS idx_maps_review_activity;
DROP INDEX IF EXISTS idx_map_tags_recent;
DROP INDEX IF EXISTS idx_map_tags_tag_discovery;
DROP INDEX IF EXISTS idx_map_tags_count_optimization;
DROP INDEX IF EXISTS idx_map_reviews_user_activity;
DROP INDEX IF EXISTS idx_map_reviews_distinct_map_names;
DROP INDEX IF EXISTS idx_maps_created_path;
DROP INDEX IF EXISTS idx_images_caption_trigram;
DROP INDEX IF EXISTS idx_map_review_images_image_review;
DROP INDEX IF EXISTS idx_map_review_images_review_image;
DROP INDEX IF EXISTS idx_map_tags_lk_tag_created;
DROP INDEX IF EXISTS idx_map_tags_map_name_lk_tag;
DROP INDEX IF EXISTS idx_map_tags_lk_tag_map_name;
