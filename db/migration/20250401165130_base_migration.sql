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
CREATE TYPE auction_status AS ENUM ('ongoing', 'finished');

CREATE TABLE auctions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	product_name text UNIQUE not null,
	product_desc text not null,
	auc_mode auction_mode not null DEFAULT 'manual',
	auc_status auction_status not null DEFAULT 'ongoing',
	
	starting_price real not null,
	target_price real,
	
	seller_id UUID references users(id)
);

-- +goose Down
DROP TABLE IF EXISTS auctions;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS auction_mode;
DROP TYPE IF EXISTS auction_status;