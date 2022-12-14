-- create "roles" table
CREATE TABLE "roles" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "phone" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "user_phone_deleted_at" to table: "users"
CREATE UNIQUE INDEX "user_phone_deleted_at" ON "users" ("phone", "deleted_at");
-- create "user_roles" table
CREATE TABLE "user_roles" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, "role_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create index "userrole_user_id_role_id" to table: "user_roles"
CREATE UNIQUE INDEX "userrole_user_id_role_id" ON "user_roles" ("user_id", "role_id");
