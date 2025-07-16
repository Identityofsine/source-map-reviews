-- +goose Up

CREATE TABLE IF NOT EXISTS user_email_whitelist (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT fk_user_email_whitelist
  FOREIGN KEY (username)
  REFERENCES user_email_whitelist(email)
  ON DELETE CASCADE;

INSERT INTO user_email_whitelist (email, description) VALUES
('identityofsine@gmail.com', 'Kevin Erdogan');

INSERT INTO user_email_whitelist (email, description) VALUES
('blank.thompson1029@protonmail.com', 'Drake');

INSERT INTO user_email_whitelist (email, description) VALUES
('maxvondoom99@gmail.com', 'Shmaxy');

-- +goose Down

DROP TABLE IF EXISTS user_email_whitelist;
