-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "citext";

CREATE TABLE users (
    id CHAR(26) PRIMARY KEY,
    name VARCHAR(255),-- display name
    username VARCHAR(255),-- username
    email CITEXT,-- email
    wallet_address CITEXT,-- wallet address
    status SMALLINT DEFAULT 0,-- 0: banned, 1: inactive, 2: active
    extra JSONB DEFAULT '{}',-- extra information
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_users_email_unique ON users(email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX idx_users_wallet_address_unique ON users(wallet_address) WHERE wallet_address IS NOT NULL;
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_extra_gin ON users USING GIN (extra);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
