-- +goose Up
--kerdogan:create_server_build_info


CREATE TABLE IF NOT EXISTS buildinfo (
	version VARCHAR(20) NOT NULL,
	commit_hash VARCHAR(128) NOT NULL,
	build_date TIMESTAMP NOT NULL,
	environment VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (version, commit_hash)
);


 -- Step 1: Add the new columns
ALTER TABLE logs
ADD COLUMN version VARCHAR(20),
ADD COLUMN commit_hash VARCHAR(128);

-- Step 2: Add the foreign key constraint
ALTER TABLE logs
ADD CONSTRAINT fk_logs_buildinfo
FOREIGN KEY (version, commit_hash)
REFERENCES buildinfo (version, commit_hash);

