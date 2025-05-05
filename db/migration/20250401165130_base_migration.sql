-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text UNIQUE not null,
    fullname text not null,
    pass_hash bytea not null,
    pass_salt bytea not null
);

CREATE TYPE auction_mode AS ENUM ('manual', 'price_met');
CREATE TYPE auction_status AS ENUM ('scheduled','ongoing', 'finished');

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_name text UNIQUE not null
);

CREATE TABLE auctions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	product_name text UNIQUE not null,
	product_desc text not null,
	auc_mode auction_mode not null DEFAULT 'manual',
	auc_status auction_status not null DEFAULT 'ongoing',
	
	starting_price real not null,
	target_price real,

    created_at timestamp not null DEFAULT now(),
	
	category_id UUID not null references categories(id),
    seller_id UUID not null references users(id)
);

CREATE VIEW auction_details AS
SELECT
    a.id,
    a.product_name,
    a.product_desc,
    a.auc_mode,
    a.auc_status,
    a.starting_price,
    a.target_price,
    a.created_at,
    u.id AS seller_id,
    u.fullname AS seller_name,
    u.email AS seller_email,
    c.id AS category_id,
    c.category_name AS category_name
FROM auctions a
JOIN users u ON a.seller_id = u.id
JOIN categories c ON a.category_id = c.id;

-- +goose Down
DROP VIEW IF EXISTS auction_details;

DROP TABLE IF EXISTS auctions;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS auction_mode;
DROP TYPE IF EXISTS auction_status;

DROP TABLE IF EXISTS categories;