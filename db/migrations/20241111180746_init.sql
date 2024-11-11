-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bytea NOT NULL,
  "username" character varying(72) NOT NULL,
  "password" character varying(72) NOT NULL,
  "email" character varying(72) NOT NULL,
  "email_verified_at" timestamp NULL,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create "collectors" table
CREATE TABLE "public"."collectors" (
  "id" bytea NOT NULL,
  "user_id" bytea NOT NULL,
  "description" character varying(1024) NULL,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "collectors_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "collections" table
CREATE TABLE "public"."collections" (
  "id" bytea NOT NULL,
  "collector_id" bytea NOT NULL,
  "title" character varying(72) NOT NULL,
  "description" character varying(1024) NULL,
  "type" character varying(72) NOT NULL,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "collections_collector_id_fkey" FOREIGN KEY ("collector_id") REFERENCES "public"."collectors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "user_tokens" table
CREATE TABLE "public"."user_tokens" (
  "user_id" bytea NOT NULL,
  "token" bytea NOT NULL,
  CONSTRAINT "user_tokens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
