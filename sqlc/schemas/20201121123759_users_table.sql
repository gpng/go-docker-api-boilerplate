-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users(
   id SERIAL PRIMARY KEY,
   name TEXT NOT NULL
);

INSERT INTO users (id, name)
VALUES (1, 'test user');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
