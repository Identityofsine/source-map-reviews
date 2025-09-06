-- +goose Up
-- +goose StatementBegin

SELECT public.create_table(
    table_name => 'lk_server_statuses',
    columns => '
      lk_server_status VARCHAR(255) NOT NULL,
      description TEXT,
      short_description VARCHAR(255) NOT NULL,
    ',
    foreign_keys => '[]',
    options => '{
        "schema": "public",
        "add_soft_delete": false,
        "add_timestamps": true,
        "primary_key": "lk_server_status",
        "comment": "Table for storing server statuses",
        "if_not_exists": true,
        "indexes": ["lk_server_status"]
    }'
);

SELECT create_table(
    table_name => 'servers',
    columns => 'server_id SERIAL,
                server_ip VARCHAR(255) NOT NULL,
                server_port INTEGER DEFAULT 27015,
                rcon_password VARCHAR(255) NOT NULL,
                server_status VARCHAR(255) NOT NULL,
    ',
    foreign_keys => '[
        {
            "column": "server_status",
            "references": "lk_server_statuses(lk_server_status)",
            "on_delete": "CASCADE"
        }
    ]',
    options => '{
        "schema": "public",
        "add_soft_delete": false,
        "add_timestamps": true,
        "primary_key": "server_id",
        "comment": "Table for storing servers",
        "if_not_exists": true,
        "indexes": ["server_ip", "server_port", "server_status"]
    }'
);

INSERT INTO lk_server_statuses (lk_server_status, description, short_description)
VALUES
('active', 'Server is active', 'active'),
('inactive', 'Server is inactive', 'inactive'),
('deleted', 'Server is deleted', 'deleted');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
