-- +goose Up
CREATE TABLE products (
	name TEXT PRIMARY KEY,
	description TEXT NOT NULL,
	price double precision NOT NULL,
	issold boolean NOT NULL DEFAULT false
);
-- +goose Down
DROP TABLE IF EXISTS products;
