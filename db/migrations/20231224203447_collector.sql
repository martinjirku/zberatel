-- migrate:up
CREATE TABLE collectors (
    id BYTEA PRIMARY KEY,
    user_id BYTEA NOT NULL,
    description VARCHAR(1024),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TRIGGER update_collector_updated_at BEFORE UPDATE ON collectors
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

INSERT INTO collectors (id, user_id) SELECT id, id FROM users;
-- migrate:down
DROP TRIGGER IF EXISTS update_collector_updated_at ON collectors;
DROP TABLE IF EXISTS collectors;
