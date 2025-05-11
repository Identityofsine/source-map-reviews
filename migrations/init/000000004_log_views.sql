-- +goose Up
--kerdogan:insert_serverity_lks

-- View for INFO logs
CREATE OR REPLACE VIEW logs_info AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.name = 'INFO';

-- View for WARNING logs
CREATE OR REPLACE VIEW logs_warning AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.name = 'WARNING';

-- View for ERROR logs
CREATE OR REPLACE VIEW logs_error AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.name = 'ERROR';

-- View for FATAL logs
CREATE OR REPLACE VIEW logs_fatal AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.name = 'FATAL';

-- View for DEBUG logs
CREATE OR REPLACE VIEW logs_debug AS
SELECT l.*
FROM logs l
JOIN severity_lks s ON l.severity = s.name
WHERE s.name = 'DEBUG';




