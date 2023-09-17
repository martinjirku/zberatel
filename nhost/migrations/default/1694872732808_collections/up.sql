CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE collections (
    id                char(27) PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMPTZ,
    title             VARCHAR(1024),
    is_blueprint      BOOLEAN NOT NULL DEFAULT FALSE,
    date_published    TIMESTAMPTZ
);

CREATE TABLE collection_items (
    id                char(27) PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMPTZ,
    title             VARCHAR(1024),
    collection_id     char(27) NOT NULL REFERENCES collections(id) ON DELETE CASCADE,
    is_blueprint      BOOLEAN NOT NULL DEFAULT FALSE,
    is_wish           BOOLEAN NOT NULL DEFAULT FALSE,
    date_published    TIMESTAMPTZ
);
CREATE INDEX idx_collection_items_collection_id ON collection_items(collection_id);

CREATE TABLE collection_item_gallery (
    id                 char(27) PRIMARY KEY,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at         TIMESTAMPTZ,
    collection_item_id char(27) NOT NULL REFERENCES collection_items(id) ON DELETE CASCADE,
    file_id            UUID NOT NULL REFERENCES storage.files(id) ON DELETE CASCADE,
    sort_order         INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX idx_collection_items_gallery_collection_item_id ON collection_item_gallery(collection_item_id);
