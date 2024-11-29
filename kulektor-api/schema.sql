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
    id          VARCHAR(512) PRIMARY KEY,
    username    VARCHAR(72) NOT NULL,
    email       VARCHAR(72) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
    is_blueprint    BOOLEAN NOT NULL DEFAULT FALSE,

    FOREIGN KEY (blueprint_id) REFERENCES collections(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
