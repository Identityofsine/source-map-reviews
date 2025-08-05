-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE public.map_review_images
ADD CONSTRAINT map_review_images_map_review_id_image_id_unique
UNIQUE (map_review_id, image_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
