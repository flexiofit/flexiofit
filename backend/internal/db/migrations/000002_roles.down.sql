-- 000001_roles.down.sql
DROP TRIGGER IF EXISTS update_roles_updated_at ON roles;
DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;
DROP TABLE IF EXISTS roles;
DROP TYPE IF EXISTS user_type;
