-- +goose Up
-- Migration: Create lk_tags table
-- Version: 000000013
-- Description: Creates the lk_tags lookup table for tag definitions

select public.create_table(
   table_name => 'lk_tags',
   columns => 'lk_tag VARCHAR(100) NOT NULL,
               description TEXT,
               ',
   options => '{
       "schema": "public",
       "add_soft_delete": false,
       "add_timestamps": true,
       "primary_key": "lk_tag",
       "comment": "Lookup table for tag definitions",
       "if_not_exists": true
   }'
);

-- +goose Down
DROP TABLE IF EXISTS lk_tags; 
