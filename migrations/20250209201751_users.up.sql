CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password CHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
    --CONSTRAINT uni_users_email UNIQUE (email)
);
-- ALTER TABLE users ADD CONSTRAINT uni_users_email UNIQUE (email);