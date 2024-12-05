DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'ksuid'
    ) THEN
        CREATE DOMAIN ksuid AS CHAR(27);
    END IF;
END $$;

CREATE TABLE users (
    id              VARCHAR(512) PRIMARY KEY,
    username        VARCHAR(72) NOT NULL,
    email           VARCHAR(72) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE collections (
    id              ksuid PRIMARY KEY,
    user_id         VARCHAR(512) NOT NULL,
    title           VARCHAR(72) NOT NULL,
    description     VARCHAR(1024),
    type            VARCHAR(72),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    blueprint_id    ksuid,

    FOREIGN KEY (blueprint_id) REFERENCES collections(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE collection_item (
    id              ksuid PRIMARY KEY,
    title           VARCHAR(1024), 
    description     TEXT,
    details         JSONB,
    meta            JSONB,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE collections_collection_item (
    collection_id           ksuid,
    collection_item_id      ksuid,
    created_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (collection_id)         REFERENCES collections(id),
    FOREIGN KEY (collection_item_id)    REFERENCES collection_item(id),

    CONSTRAINT unique_collections_collection_item UNIQUE (collection_id, collection_item_id)
);

CREATE TABLE blueprints (
    id              ksuid PRIMARY KEY,
    user_id         VARCHAR(512),
    title           VARCHAR(72) NOT NULL,
    description     VARCHAR(1024),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blueprint_item (
    id              ksuid PRIMARY KEY,
    title           VARCHAR(1024),
    description     TEXT,
    product_code    VARCHAR(512),
    details         JSONB,
    meta            JSONB,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blueprints_blueprint_item (
    blueprints_id       ksuid,
    blueprint_item_id   ksuid,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (blueprints_id)         REFERENCES blueprints(id),
    FOREIGN KEY (blueprint_item_id)     REFERENCES blueprint_item(id),

    CONSTRAINT unique_blueprints_blueprint_item UNIQUE (blueprints_id, blueprint_item_id)
);

CREATE UNIQUE INDEX idx_blueprints_blueprint_item_unique ON blueprints_blueprint_item (blueprints_id, blueprint_item_id);

CREATE TABLE documents (
    id              ksuid PRIMARY KEY,
    table_name      VARCHAR(72) NOT NULL,
    table_id        ksuid,
    title           VARCHAR(512) NOT NULL,
    type            VARCHAR(72) NOT NULL,
    location        VARCHAR(512) NOT NULL,
    user_id         VARCHAR(512),

    CONSTRAINT valid_table CHECK (table_name IN ('blueprints', 'blueprint_item', 'collections', 'collection_item'))
);