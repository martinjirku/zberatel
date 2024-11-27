DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'ksuid'
    ) THEN
        CREATE DOMAIN ksuid AS BYTEA;
    END IF;
END $$;

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE users (
    id VARCHAR(512) PRIMARY KEY,
    username VARCHAR(72) NOT NULL,
    email VARCHAR(72) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_user_updated_at BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE collections (
    id ksuid PRIMARY KEY,
    user_id VARCHAR(512) NOT NULL,
    title VARCHAR(72) NOT NULL,
    description VARCHAR(1024),
    type VARCHAR(72) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE OR REPLACE TRIGGER update_collection_updated_at BEFORE UPDATE ON collections
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();