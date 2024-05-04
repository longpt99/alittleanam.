-- +goose Up
-- +goose StatementBegin

-- Create user status enum
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'users_status_enum') THEN
        CREATE TYPE users_status_enum AS ENUM ('active', 'inactive');
    END IF;
END$$;

CREATE UNIQUE INDEX IF NOT EXISTS users_email_unique 
  ON users (email) 
  WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS users
  (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email         TEXT NOT NULL,
    username      TEXT NULL,
    password      TEXT NOT NULL,
    first_name    TEXT NOT NULL,
    last_name     TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    image_url     TEXT NULL,
    story         TEXT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ NULL
    status        status users_status_enum NOT NULL DEFAULT 'inactive',
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS users_status_enum;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd