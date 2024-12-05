-- Create "blueprint_item" table
CREATE TABLE "blueprint_item" (
 "id" ksuid NOT NULL,
 "title" character varying(1024) NULL,
 "description" text NULL,
 "product_code" character varying(512) NULL,
 "details" jsonb NULL,
 "meta" jsonb NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY ("id")
);
-- Create "documents" table
CREATE TABLE "documents" (
 "id" ksuid NOT NULL,
 "table_name" character varying(72) NOT NULL,
 "table_id" ksuid NULL,
 "title" character varying(512) NOT NULL,
 "type" character varying(72) NOT NULL,
 "location" character varying(512) NOT NULL,
 "user_id" character varying(512) NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "valid_table" CHECK ((table_name)::text = ANY ((ARRAY['blueprints'::character varying, 'blueprint_item'::character varying, 'collections'::character varying, 'collection_item'::character varying])::text[]))
);
-- Create "blueprints" table
CREATE TABLE "blueprints" (
 "id" ksuid NOT NULL,
 "user_id" character varying(512) NULL,
 "title" character varying(72) NOT NULL,
 "description" character varying(1024) NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY ("id")
);
-- Create "blueprints_blueprint_item" table
CREATE TABLE "blueprints_blueprint_item" (
 "blueprints_id" ksuid NULL,
 "blueprint_item_id" ksuid NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 CONSTRAINT "unique_blueprints_blueprint_item" UNIQUE ("blueprints_id", "blueprint_item_id"),
 CONSTRAINT "blueprints_blueprint_item_blueprint_item_id_fkey" FOREIGN KEY ("blueprint_item_id") REFERENCES "blueprint_item" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT "blueprints_blueprint_item_blueprints_id_fkey" FOREIGN KEY ("blueprints_id") REFERENCES "blueprints" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_blueprints_blueprint_item_unique" to table: "blueprints_blueprint_item"
CREATE UNIQUE INDEX "idx_blueprints_blueprint_item_unique" ON "blueprints_blueprint_item" ("blueprints_id", "blueprint_item_id");
-- Modify "collections" table
ALTER TABLE "collections" DROP COLUMN "is_blueprint";
-- Create "collection_item" table
CREATE TABLE "collection_item" (
 "id" ksuid NOT NULL,
 "title" character varying(1024) NULL,
 "description" text NULL,
 "details" jsonb NULL,
 "meta" jsonb NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY ("id")
);
-- Create "collections_collection_item" table
CREATE TABLE "collections_collection_item" (
 "collection_id" ksuid NULL,
 "collection_item_id" ksuid NULL,
 "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 CONSTRAINT "unique_collections_collection_item" UNIQUE ("collection_id", "collection_item_id"),
 CONSTRAINT "collections_collection_item_collection_id_fkey" FOREIGN KEY ("collection_id") REFERENCES "collections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT "collections_collection_item_collection_item_id_fkey" FOREIGN KEY ("collection_item_id") REFERENCES "collection_item" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
