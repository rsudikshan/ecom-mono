-- +migrate Up
ALTER TABLE users 
ADD COLUMN role TEXT DEFAULT 'ROLE_USER' NOT NULL;

-- +migrate Down
ALTER TABLE users
DROP COLUMN role;