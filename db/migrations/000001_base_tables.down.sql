
-- Drop base tables, should they already exists
DROP TABLE IF EXISTS snacks;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

-- Drop the enums, should they already exist
DROP TYPE IF EXISTS rinks_enum;
DROP TYPE IF EXISTS levels_enum;
DROP TYPE IF EXISTS snack_rankings_enum;


-- No need to remove the functions
