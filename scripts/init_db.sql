-- DROP CREATE DB
DROP DATABASE IF EXISTS basic_db;
CREATE DATABASE basic_db;

-- DROP CREATE TABLE
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `email` varchar NOT NULL,
    `name` varchar NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at bigint,
    updated_at bigint
);

-- DROP CREATE TABLE
DROP TABLE IF EXISTS contents;
CREATE TABLE contents (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `name` varchar NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at bigint,
    updated_at bigint,
    user_id uuid NOT NULL REFERENCES users(id)
);
