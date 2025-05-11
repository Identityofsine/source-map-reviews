-- +goose Up
--kerdogan:insert_serverity_lks

-- View for INFO logs
CREATE OR REPLACE VIEW logs_info AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.level = 0;

-- View for WARNING logs
CREATE OR REPLACE VIEW logs_warning AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.level = 1;

-- View for ERROR logs
CREATE OR REPLACE VIEW logs_error AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.level = 2;

-- View for FATAL logs
CREATE OR REPLACE VIEW logs_fatal AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.level = 3;
