-- +goose Up
-- Migration: Create creator table
-- Version: 000000010
-- Description: Creates the creator table with auto-increment primary key

select public.create_table(
   table_name => 'creator',
   columns => 'creator_id SERIAL,
               creator_name VARCHAR(255) NOT NULL,
               ',
   options => '{
       "schema": "public",
       "add_id": false,
       "add_soft_delete": false,
       "add_timestamps": true,
       "primary_key": "creator_id",
       "comment": "Table for storing map creators",
       "if_not_exists": true,
       "indexes": ["creator_name"]
   }'
);

-- +goose Down
DROP TABLE IF EXISTS creator; 
