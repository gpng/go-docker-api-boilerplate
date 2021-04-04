-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users(
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   email TEXT UNIQUE NOT NULL,
   password TEXT NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
