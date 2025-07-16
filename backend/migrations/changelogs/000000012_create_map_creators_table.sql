-- +goose Up
-- Migration: Create map_creators table
-- Version: 000000012
-- Description: Creates the map_creators junction table linking map_infos to creators

select public.create_table(
   table_name => 'map_creators',
   columns => 'map_name VARCHAR(255) NOT NULL,
               creator_id INTEGER NOT NULL,
               ',
   foreign_keys => '[
       {
           "column": "map_name",
           "references": "map_infos(map_name)",
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
       "add_timestamps": false,
       "comment": "Junction table for map creators",
       "if_not_exists": true
   }'
);

-- +goose Down
DROP TABLE IF EXISTS map_creators; 