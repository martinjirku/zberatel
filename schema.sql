CREATE DOMAIN ksuid AS BYTEA;

CREATE TABLE users (
    -- reconsider the BYTEA type for ksuid
    id ksuid PRIMARY KEY,
    username VARCHAR(72) NOT NULL,
    password VARCHAR(72) NOT NULL,
    email VARCHAR(72) NOT NULL,
    email_verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE user_tokens (
    -- reconsider the BYTEA vs VARCHAR(27) type for ksuid
    user_id ksuid NOT NULL,
    token ksuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE collectors (
    id ksuid PRIMARY KEY,
    user_id ksuid NOT NULL,
    description VARCHAR(1024),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TRIGGER update_collector_updated_at BEFORE UPDATE ON collectors
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE collections (
    id ksuid PRIMARY KEY,
    user_id ksuid NOT NULL,
    title VARCHAR(72) NOT NULL,
    description VARCHAR(1024),
    type VARCHAR(72) NOT NULL, -- TODO: extract to a separate table/enum/whatever
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TRIGGER update_collection_updated_at BEFORE UPDATE ON collections
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();