CREATE DATABASE app;

CREATE USER docker WITH PASSWORD 'docker';

GRANT ALL PRIVILEGES ON DATABASE app TO docker;

-- Grant the necessary permissions to the docker user
GRANT USAGE, CREATE ON SCHEMA public TO docker;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO docker;

-- Ensure future tables created by any user will be accessible
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO docker;
