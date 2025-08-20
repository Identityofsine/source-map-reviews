-- +goose Up
-- Migration: Add map search performance indexes
-- Description: Creates indexes to optimize map name search queries with ILIKE patterns and text search

-- Enable pg_trgm extension for trigram indexes (for ILIKE optimization)
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Trigram index for case-insensitive map name searches (ILIKE queries)
-- This dramatically improves performance for substring, prefix, and suffix matching
CREATE INDEX IF NOT EXISTS idx_maps_map_name_trigram 
ON maps USING gin (map_name gin_trgm_ops);

-- Composite index for maps with timestamps for sorting recent maps
CREATE INDEX IF NOT EXISTS idx_maps_created_at_map_name 
ON maps (created_at DESC, map_name);

-- Partial index for active maps
CREATE INDEX IF NOT EXISTS idx_maps_active_map_name 
ON maps (map_name) WHERE created_at IS NOT NULL;

-- +goose Down
-- Drop indexes in reverse order
DROP INDEX IF EXISTS idx_maps_active_map_name;
DROP INDEX IF EXISTS idx_maps_created_at_map_name;
DROP INDEX IF EXISTS idx_maps_map_name_trigram;
