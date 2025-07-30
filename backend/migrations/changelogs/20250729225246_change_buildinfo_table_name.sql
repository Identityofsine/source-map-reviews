-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE buildinfo RENAME TO build_infos;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
