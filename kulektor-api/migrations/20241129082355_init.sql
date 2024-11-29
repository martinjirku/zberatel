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
-- Create "users" table
CREATE TABLE "users" (
 "id" character varying(512) NOT NULL,
 "username" character varying(72) NOT NULL,
 "email" character varying(72) NOT NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY ("id")
);
-- Create "collections" table
CREATE TABLE "collections" (
 "id" ksuid NOT NULL,
 "user_id" character varying(512) NOT NULL,
 "title" character varying(72) NOT NULL,
 "description" character varying(1024) NULL,
 "type" character varying(72) NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 "blueprint_id" ksuid NULL,
 "is_blueprint" boolean NOT NULL DEFAULT false,
 PRIMARY KEY ("id"),
 CONSTRAINT "collections_blueprint_id_fkey" FOREIGN KEY ("blueprint_id") REFERENCES "collections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT "collections_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
