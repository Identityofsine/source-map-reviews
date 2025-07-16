-- +goose Up
--kerdogan:authtokens

CREATE TABLE IF NOT EXISTS authtokens (
	id SERIAL PRIMARY KEY,
  user_id integer NOT NULL,
  access_token varchar(255) NOT NULL,
  refresh_token varchar(255) NOT NULL,
	expires_at timestamp NOT NULL,
	refreshed_at timestamp DEFAULT CURRENT_TIMESTAMP,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
