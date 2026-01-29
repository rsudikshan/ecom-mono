-- +migrate Up
CREATE TABLE users(
    id CHAR(27) PRIMARY KEY,
    -- firstname TEXT NOT NULL,
    -- lastname TEXT NOT NULL,
    phone TEXT,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL,
    password TEXT,
    email_verified BOOLEAN DEFAULT FALSE NOT NULL,

    created_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL,

    created_by_id CHAR(27) REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id CHAR(27) REFERENCES users(id) ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;