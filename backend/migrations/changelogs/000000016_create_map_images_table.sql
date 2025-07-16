-- +goose Up
-- Migration: Create map_images table
-- Version: 000000016
-- Description: Creates the map_images junction table linking maps to images

select public.create_table(
   table_name => 'map_images',
   columns => 'map_image_id SERIAL,
               map_name VARCHAR(255) NOT NULL,
               image_id INTEGER NOT NULL,
               ',
   foreign_keys => '[
       {
           "column": "map_name",
           "references": "maps(map_name)",
           "on_delete": "CASCADE"
       },
       {
           "column": "image_id",
           "references": "images(image_id)",
           "on_delete": "CASCADE"
       }
   ]',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": false,
       "primary_key": "map_image_id",
       "comment": "Junction table linking maps to images",
       "if_not_exists": true,
       "indexes": ["map_name", "image_id"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS map_images; 