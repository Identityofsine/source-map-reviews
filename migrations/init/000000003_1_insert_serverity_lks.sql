-- +goose Up
--kerdogan:insert_serverity_lks

CALL public.create_severity_lk('INFO', 'Informational messages', 0);
CALL public.create_severity_lk('WARNING', 'Warning messages', 1);
CALL public.create_severity_lk('ERROR', 'Error messages', 2);
CALL public.create_severity_lk('FATAL', 'Fatal error messages (Crashes)', 3);
CALL public.create_severity_lk('DEBUG', 'Debug messages', 4);


