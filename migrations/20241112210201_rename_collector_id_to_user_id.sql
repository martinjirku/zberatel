-- Modify "collections" table
ALTER TABLE "public"."collections" DROP COLUMN "collector_id", ADD COLUMN "user_id" bytea NOT NULL, ADD
 CONSTRAINT "collections_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
