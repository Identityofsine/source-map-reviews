-- +goose Up
-- Migration: Add map search performance indexes
-- Description: Creates indexes to optimize map name search queries with ILIKE patterns and text search

-- Enable pg_trgm extension for trigram indexes (for ILIKE optimization)
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Trigram index for case-insensitive map name searches (ILIKE queries)
-- This dramatically improves performance for substring, prefix, and suffix matching
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_map_name_trigram 
ON maps USING gin (map_name gin_trgm_ops);

-- Case-insensitive index for exact and prefix matching
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_map_name_lower 
ON maps (lower(map_name));

-- Index for map_path searches (in case we need to search by file path)
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_map_path_lower 
ON maps (lower(map_path));

-- Composite index for maps with timestamps for sorting recent maps
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_created_at_map_name 
ON maps (created_at DESC, map_name);

-- Partial index for active maps (if we ever add soft delete)
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_maps_active_map_name 
ON maps (map_name) WHERE created_at IS NOT NULL;

-- +goose Down
-- Drop indexes in reverse order
DROP INDEX IF EXISTS idx_maps_active_map_name;
DROP INDEX IF EXISTS idx_maps_created_at_map_name;
DROP INDEX IF EXISTS idx_maps_map_path_lower;
DROP INDEX IF EXISTS idx_maps_map_name_lower;
DROP INDEX IF EXISTS idx_maps_map_name_trigram;
