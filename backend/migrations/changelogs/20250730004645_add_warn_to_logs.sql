-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

SELECT public.create_severity_lk('WARN', 'Warning severity level', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
