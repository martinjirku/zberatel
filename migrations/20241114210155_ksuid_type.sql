-- Modify "collections" table
ALTER TABLE "collections" ALTER COLUMN "id" TYPE ksuid, ALTER COLUMN "user_id" TYPE ksuid;
-- Modify "collectors" table
ALTER TABLE "collectors" ALTER COLUMN "id" TYPE ksuid, ALTER COLUMN "user_id" TYPE ksuid;
-- Modify "user_tokens" table
ALTER TABLE "user_tokens" ALTER COLUMN "user_id" TYPE ksuid, ALTER COLUMN "token" TYPE ksuid;
-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "id" TYPE ksuid;
