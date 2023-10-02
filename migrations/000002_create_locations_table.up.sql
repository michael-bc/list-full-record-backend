CREATE TABLE IF NOT EXISTS locations
(
    id         SERIAL PRIMARY KEY,
    city       VARCHAR   NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);