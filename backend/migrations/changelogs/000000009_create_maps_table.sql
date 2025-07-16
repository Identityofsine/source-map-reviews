-- +goose Up
-- Migration: Create maps table
-- Version: 000000009
-- Description: Creates the maps table with map_name as primary key and map_path

select public.create_table(
   table_name => 'maps',
   columns => 'map_name VARCHAR(255) NOT NULL,
               map_path VARCHAR(500) NOT NULL,
               ',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": true,
       "primary_key": "map_name",
       "comment": "Table for storing map files with their names and paths",
       "if_not_exists": true,
       "indexes": ["map_path"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS maps; 
