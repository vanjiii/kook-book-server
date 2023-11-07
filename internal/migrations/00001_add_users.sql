-- +goose Up
-- Postgres module that provides cryptographic functions used for passwords
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT PUBLIC.gen_random_uuid(),
        email TEXT UNIQUE NOT NULL,
        password BYTEA NOT NULL,
        name TEXT,
        avatar TEXT,

	created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	deleted_at TIMESTAMP WITH TIME ZONE
);

-- +goose Down
DROP TABLE users;
