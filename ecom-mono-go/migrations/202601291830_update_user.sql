-- +migrate Up
ALTER TABLE users
ADD COLUMN password_reset_at TIMESTAMPTZ DEFAULT NULL;

-- +migrate Down
ALTER TABLE users
DROP COLUMN password_reset_at;