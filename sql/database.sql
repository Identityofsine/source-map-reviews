-- Create user and database
CREATE USER docker WITH PASSWORD 'docker';
CREATE DATABASE app OWNER docker;

-- Connect to app database
\c app

-- Optional: Make sure docker owns the schema
ALTER SCHEMA public OWNER TO docker;

-- Grant privileges on schema and future objects
GRANT USAGE, CREATE ON SCHEMA public TO docker;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO docker;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO docker;

-- Ensure future tables/sequences/etc. are accessible
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO docker;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO docker;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON FUNCTIONS TO docker;
