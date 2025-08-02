-- +goose Up
-- Migration: Create map_reviews table
-- Version: 000000017
-- Description: Creates the map_reviews table for storing user reviews of maps

select public.create_table(
   table_name => 'map_reviews',
   columns => 'map_review_id SERIAL AUTO_INCREMENT,
               map_name VARCHAR(255) NOT NULL,
               reviewer INTEGER NOT NULL,
               stars INTEGER CHECK (stars >= 1 AND stars <= 5),
               review TEXT,
               ',
   foreign_keys => '[
       {
           "column": "map_name",
           "references": "maps(map_name)",
           "on_delete": "CASCADE"
       },
       {
           "column": "reviewer",
           "references": "users(id)",
           "on_delete": "CASCADE"
       }
   ]',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": true,
       "primary_key": "map_review_id",
       "comment": "Table for storing user reviews of maps",
       "if_not_exists": true,
       "indexes": ["map_name", "reviewer", "stars"]
   }'
);


-- +goose Down
DROP TABLE IF EXISTS map_reviews; 
