-- +goose Up
-- Migration: Create map_tags table
-- Version: 000000014
-- Description: Creates the map_tags junction table linking maps to tags with unique constraint

select public.create_table(
   table_name => 'map_tags',
   columns => 'lk_tag VARCHAR(100) NOT NULL,
               map_name VARCHAR(255) NOT NULL,
               ',
   foreign_keys => '[
       {
           "column": "lk_tag",
           "references": "lk_tags(lk_tag)",
           "on_delete": "CASCADE"
       },
       {
           "column": "map_name",
           "references": "maps(map_name)",
           "on_delete": "CASCADE"
       }
   ]',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": false,
       "unique_constraints": ["lk_tag, map_name"],
       "comment": "Junction table linking maps to tags",
       "if_not_exists": true,
       "indexes": ["map_name"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS map_tags; 