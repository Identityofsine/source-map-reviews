-- +goose Up
-- Migration: Create map_infos table
-- Version: 000000011
-- Description: Creates the map_infos table with foreign key references to maps and creator

select public.create_table(
   table_name => 'map_infos',
   columns => 'map_name VARCHAR(255) NOT NULL,
               creator_id INTEGER NOT NULL,
               description TEXT,
               ',
   foreign_keys => '[
       {
           "column": "map_name",
           "references": "maps(map_name)",
           "on_delete": "CASCADE"
       },
       {
           "column": "creator_id",
           "references": "creator(creator_id)",
           "on_delete": "CASCADE"
       }
   ]',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": true,
       "primary_key": "map_name",
       "comment": "Table for storing map information and metadata",
       "if_not_exists": true,
       "indexes": ["creator_id"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS map_infos; 