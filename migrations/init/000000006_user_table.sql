-- +goose Up
--kerdogan:authtokens


CREATE TABLE IF NOT EXISTS authentication_method_lks (
	name VARCHAR(20) PRIMARY KEY,
	description VARCHAR(50) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username varchar(255) NOT NULL,
	password varchar(255),
	authentication_method varchar(20) NOT NULL,
	verified boolean DEFAULT false,
	FOREIGN KEY (authentication_method) REFERENCES authentication_method_lks(name)
);
