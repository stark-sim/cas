-- reverse: create index "userrole_user_id_role_id" to table: "user_roles"
DROP INDEX "userrole_user_id_role_id";
-- reverse: create "user_roles" table
DROP TABLE "user_roles";
-- reverse: create index "user_phone_deleted_at" to table: "users"
DROP INDEX "user_phone_deleted_at";
-- reverse: create "users" table
DROP TABLE "users";
-- reverse: create "roles" table
DROP TABLE "roles";
