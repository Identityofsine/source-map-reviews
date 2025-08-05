-- +goose Up
-- +goose StatementBegin

DROP TABLE IF EXISTS map_review_images;
DROP TABLE IF EXISTS map_reviews_images;

SELECT public.create_table(
    table_name => 'map_review_images',
    columns => 'map_review_image_id SERIAL,
                map_review_id INTEGER NOT NULL,
                image_id INTEGER NOT NULL,',
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
        "primary_key": "map_review_image_id",
        "if_not_exists": true,
        "indexes": ["map_review_id", "image_id"]
    }'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
