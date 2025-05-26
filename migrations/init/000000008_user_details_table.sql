-- +goose Up
--kerdogan:authtokens

CREATE TABLE IF NOT EXISTS user_details (
	id SERIAL PRIMARY KEY,
  user_id integer NOT NULL,
	email VARCHAR(255),
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	date_of_birth DATE,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_oauth_tokens (
	user_id integer NOT NULL,
	access_token VARCHAR(255) NOT NULL,
	refresh_token VARCHAR(255) NOT NULL,
	source VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	expires_at TIMESTAMP,
	PRIMARY KEY (user_id, source),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (source) REFERENCES authentication_method_lks(name) ON DELETE CASCADE
);
