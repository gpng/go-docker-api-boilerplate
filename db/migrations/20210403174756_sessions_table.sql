-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS sessions(
   user_id UUID PRIMARY KEY REFERENCES users(id),
   refresh_token TEXT UNIQUE NOT NULL,
   invalidated BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS sessions;
