CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users (
    id UUID primary key not null default uuid_generate_v1(),
    username varchar(255) unique not null,
    email varchar(255) unique not null,
    password text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);