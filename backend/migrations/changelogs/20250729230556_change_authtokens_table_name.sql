-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE authtokens RENAME TO auth_tokens;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
