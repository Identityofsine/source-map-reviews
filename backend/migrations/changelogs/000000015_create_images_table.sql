-- +goose Up
-- Migration: Create images table
-- Version: 000000015
-- Description: Creates the images table for storing image metadata

select public.create_table(
   table_name => 'images',
   columns => 'image_id SERIAL,
               image_path VARCHAR(500) NOT NULL,
               caption TEXT,
               ',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": false,
       "primary_key": "image_id",
       "comment": "Table for storing image metadata",
       "if_not_exists": true,
       "indexes": ["image_path"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS images; 