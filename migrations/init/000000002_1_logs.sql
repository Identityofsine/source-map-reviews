-- +goose Up

CREATE TABLE IF NOT EXISTS severity_lks (
	name VARCHAR(20) PRIMARY KEY,
	description VARCHAR(50) NOT NULL,
	level INTEGER NOT NULL, -- 0 = info, 1 = warning, 2 = error, 3 = fatal
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS logs (
	id SERIAL PRIMARY KEY,
	severity VARCHAR(20) NOT NULL,
	message TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (severity) REFERENCES severity_lks(name)
);

-- +goose Down

DROP TABLE IF EXISTS logs;
DROP TABLE IF EXISTS severity_lks;
