-- Modify "collections" table
ALTER TABLE "collections" ADD COLUMN "blueprint_id" ksuid NULL, ADD COLUMN "is_blueprint" boolean NOT NULL DEFAULT false, ADD
 CONSTRAINT "collections_blueprint_id_fkey" FOREIGN KEY ("blueprint_id") REFERENCES "collections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
