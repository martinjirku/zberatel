-- migrate:up
CREATE TABLE users (
    -- reconsider the BYTEA type for ksuid
    id BYTEA PRIMARY KEY,
    username VARCHAR(72) NOT NULL,
    password VARCHAR(72) NOT NULL,
    email VARCHAR(72) NOT NULL,
    email_verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- handling the email verification token
CREATE TABLE user_tokens (
    -- reconsider the BYTEA vs VARCHAR(27) type for ksuid
    user_id BYTEA NOT NULL,
    token BYTEA NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

-- migrate:down
DROP TRIGGER IF EXISTS update_user_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS user_tokens;
DROP TABLE IF EXISTS users;
