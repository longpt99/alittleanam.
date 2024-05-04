-- +goose Up
-- +goose StatementBegin
-- #
-- # Create function to update updated_at timestamp if changed values on update
-- #
CREATE OR REPLACE FUNCTION before_update_updated_at() RETURNS trigger AS
$BODY$
BEGIN
    IF row(NEW.*::text) IS DISTINCT FROM row(OLD.*::text) THEN
        NEW.updated_at = now();
    END IF;

    RETURN NEW;
END;
$BODY$

LANGUAGE plpgsql;

-- #
-- # Apply before_update_updated_at function to all tables as trigger
-- #
DO $BODY$
DECLARE t text;
BEGIN
    FOR t IN
        SELECT table_name
        FROM information_schema.columns
        WHERE (
            column_name = 'updated_at'
            AND (
                SELECT 1
                FROM information_schema.triggers
                WHERE trigger_name = 'before_update_updated_at_' || table_name
            ) IS NULL
        )
    LOOP
        EXECUTE format('
            CREATE TRIGGER before_update_updated_at_%s
            BEFORE UPDATE ON %I
            FOR EACH ROW EXECUTE PROCEDURE before_update_updated_at();
        ', t, t);
    END loop;
END;
$BODY$

LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS before_update_updated_at() CASCADE;
-- +goose StatementEnd
