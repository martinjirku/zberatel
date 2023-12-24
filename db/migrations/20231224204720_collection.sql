-- migrate:up
CREATE TABLE collections (
    id BYTEA PRIMARY KEY,
    collector_id BYTEA NOT NULL,
    name VARCHAR(72) NOT NULL,
    description VARCHAR(1024),
    type VARCHAR(72) NOT NULL, -- TODO: extract to a separate table/enum/whatever
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (collector_id) REFERENCES collectors(id)
);

CREATE TRIGGER update_collection_updated_at BEFORE UPDATE ON collections
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- migrate:down
DROP TRIGGER IF EXISTS update_collection_updated_at ON collections;
DROP TABLE IF EXISTS collections;
