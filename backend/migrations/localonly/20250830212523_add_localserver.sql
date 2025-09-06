-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

INSERT INTO servers (server_ip, server_port, rcon_password, server_status)
VALUES
('10.2.2.2', 27015, 'mimi', 'active');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
