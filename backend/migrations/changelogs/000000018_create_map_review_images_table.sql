-- +goose Up
-- Migration: Create map_review_images table
-- Version: 000000018
-- Description: Creates the map_review_images junction table linking reviews to images

select public.create_table(
   table_name => 'map_review_images',
   columns => 'map_review_id INTEGER NOT NULL,
               image_id INTEGER NOT NULL,
               ',
   foreign_keys => '[
       {
           "column": "map_review_id",
           "references": "map_reviews(map_review_id)",
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
       "primary_key": "map_review_id",
       "comment": "Junction table linking map reviews to images",
       "if_not_exists": true,
       "indexes": ["map_review_id", "image_id"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS map_review_images; 