-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE public.create_severity_lk(
    severity_name VARCHAR,
    severity_description VARCHAR,
    severity_level INTEGER
)
LANGUAGE plpgsql
AS $$
DECLARE
    severity_count INTEGER;
BEGIN
    IF severity_name IS NULL OR severity_description IS NULL OR severity_level IS NULL THEN
        RAISE EXCEPTION 'Severity name, description, and level cannot be NULL';
    END IF;

    IF severity_level < 0 THEN
        RAISE EXCEPTION 'Severity level must be above 0';
    END IF;

    SELECT COUNT(*) INTO severity_count FROM public.severity_lks WHERE name = severity_name;

    IF severity_count > 0 THEN
        RAISE EXCEPTION 'Severity already exists';
    END IF;

    INSERT INTO public.severity_lks (name, description, level)
    VALUES (severity_name, severity_description, severity_level);
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS public.create_severity_lk(VARCHAR, VARCHAR, INTEGER);
-- +goose StatementEnd

