-- +goose Up
-- This is executed when the migration is applied
-- Add the new column without a default value
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE;

-- Set the default value using an UPDATE statement
UPDATE users SET api_key = encode(sha256(random()::text::bytea), 'hex') WHERE api_key IS NULL;

-- Make the column NOT NULL after setting the default value
ALTER TABLE users ALTER COLUMN api_key SET NOT NULL;

-- +goose Down
-- This is executed when the migration is rolled back
ALTER TABLE users DROP COLUMN api_key;
