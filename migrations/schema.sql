CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "categories" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"description" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "items" (
"id" TEXT PRIMARY KEY,
"name" text NOT NULL,
"description" text NOT NULL,
"price" float NOT NULL,
"category_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
