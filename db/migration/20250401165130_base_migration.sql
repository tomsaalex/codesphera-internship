-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text UNIQUE not null,
    fullname text not null,
    pass_hash bytea not null,
    pass_salt bytea not null
);

-- +goose Down
DROP TABLE IF EXISTS users;

