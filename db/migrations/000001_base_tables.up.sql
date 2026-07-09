
-- Create the initial tables


-- this is the enum to hold the rinks
-- REMINDER - you can add more values later, but you can't remove values
-- this could be in its own table if we anticipated it needing to change more frequently or needing more details
CREATE TYPE rinks_enum AS ENUM ('BAIREL', 'UPMC');
COMMENT ON TYPE rinks_enum IS 'The list of rinks available when defining a team in the teams table.';


-- this is the enum to hold the levels
-- REMINDER - you can add more values later, but you can't remove values
-- this could be in its own table if we anticipated it needing to change more frequently or needing more details
CREATE TYPE levels_enum AS ENUM ('D5', 'D4', 'D3');
COMMENT ON TYPE levels_enum IS 'The list of levels available when defining a team in the teams table';


-- this is the function to update the updated column on tables
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';


-- this a function to force uppercase on name columns to ensure data consistency for Snacks, Ingredients & Teams
CREATE OR REPLACE FUNCTION force_uppercase_name()
RETURNS TRIGGER AS $$
BEGIN
    NEW.name := UPPER(NEW.name);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


/*
 * teams table
 *
 */
CREATE TABLE teams (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL UNIQUE,
    rink rinks_enum NOT NULL,
    level levels_enum NOT NULL,
    primary_color TEXT NOT NULL,
    secondary_color TEXT NOT NULL,
    ternary_color TEXT NOT NULL DEFAULT '',
    logo_url TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE teams IS 'The full details of a beer league hockey team that a user might be a member of';
COMMENT ON COLUMN teams.id IS 'The unique team identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN teams.name IS 'The name of the team, cannot be null and must be unique';
COMMENT ON COLUMN teams.rink IS 'The primary rink that this team plays at, selected from the rinks_enum';
COMMENT ON COLUMN teams.level IS 'The level of the team, selected from the levels_enum';
COMMENT ON COLUMN teams.primary_color IS 'The primary color of the team, stored as the hex color code, cannot be null';
COMMENT ON COLUMN teams.secondary_color IS 'The secondary color of the team, stored as the hex color code, cannot be null';
COMMENT ON COLUMN teams.ternary_color IS 'The third color of the team, stored as the hex color code, defaults to "" if no third color exists';
COMMENT ON COLUMN teams.logo_url IS 'The path to the team logo, defaults to "" if no image url exists';
COMMENT ON COLUMN teams.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN teams.updated_at IS 'The time this row was last updated, UTC time';

-- add the triggers to teams
CREATE TRIGGER update_teams_updated_at
BEFORE UPDATE ON teams FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_teams_updated_at ON teams IS 'Update the updated_at timestamp column for the teams table';

CREATE TRIGGER force_uppercase_team_name
BEFORE INSERT OR UPDATE ON teams FOR EACH ROW EXECUTE FUNCTION force_uppercase_name();
COMMENT ON TRIGGER force_uppercase_team_name ON teams IS 'Prior to insert or update, force the name column to be all uppercase for data consistency';

/*
 * users table
 * Team membership is recorded in the team_membership table
 * Allergies are recorded in the user_allergies table
 */
CREATE TABLE users (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE users IS 'The full details of a beer league hockey team player';
COMMENT ON COLUMN users.id IS 'The unique player identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN users.first_name IS 'The first name of the user, cannot be null';
COMMENT ON COLUMN users.last_name IS 'The last name of the user, cannot be null';
COMMENT ON COLUMN users.email IS 'The email address of the user, cannot be null and must be unique';
COMMENT ON COLUMN users.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN users.updated_at IS 'The time this row was last updated, UTC time';

-- add the trigger to users
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_users_updated_at ON users IS 'Update the updated_at timestamp column for the users table';






/*
 * ingredients table
 * this is a table because we will likely need to add more values over time
 */
CREATE TABLE ingredients (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE ingredients IS 'Possible ingredients for snacks or items that a user might be allergic to';
COMMENT ON COLUMN ingredients.id IS 'The unique ingredient identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN ingredients.name IS 'The basic name of the ingredient, cannot be null and must be unique';
COMMENT ON COLUMN ingredients.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN ingredients.updated_at IS 'The time this row was last updated, UTC time';

-- add the triggers to ingredients
CREATE TRIGGER update_ingredients_updated_at
BEFORE UPDATE ON ingredients FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_ingredients_updated_at ON ingredients IS 'Update the updated_at timestamp column for the allingredientsergies table';

CREATE TRIGGER force_uppercase_ingredient_name
BEFORE INSERT OR UPDATE ON ingredients FOR EACH ROW EXECUTE FUNCTION force_uppercase_name();
COMMENT ON TRIGGER force_uppercase_ingredient_name ON ingredients IS 'Prior to insert or update, force the name column to be all uppercase for data consistency';



/*
 * snacks table
 * Ingredients are recorded in the snack_ingredients table
 */
CREATE TABLE snacks (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL UNIQUE,
    sweet BOOLEAN NOT NULL,
    savory BOOLEAN NOT NULL,
    difficulty INT NOT NULL,
    recipe_url TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE snacks IS 'Snacks and their descriptive qualities';
COMMENT ON COLUMN snacks.id IS 'The unique snack identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN snacks.name IS 'The name of the snack, not null and must be unique';
COMMENT ON COLUMN snacks.sweet IS 'Whether the snack is considered sweet, true/flase, cannot be null';
COMMENT ON COLUMN snacks.savory IS 'Whether the snack is considered savory, true/false, cannot be null';
COMMENT ON COLUMN snacks.recipe_url IS 'The url of the recipe, defaults to "" if no url exists';
COMMENT ON COLUMN snacks.difficulty IS 'Arbitrary rating by Britni on the difficulty of the recipe - includes time to prepare and ingredients required. Scale of 1 (easy) to 10 (monstrous)';
COMMENT ON COLUMN snacks.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN snacks.updated_at IS 'The time this row was last updated, UTC time';

-- add the triggers to snacks
CREATE TRIGGER update_snacks_updated_at
BEFORE UPDATE ON snacks FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_snacks_updated_at ON snacks IS 'Update the updated_at timestamp column for the snacks table';

CREATE TRIGGER force_uppercase_snack_name
BEFORE INSERT OR UPDATE ON snacks FOR EACH ROW EXECUTE FUNCTION force_uppercase_name();
COMMENT ON TRIGGER force_uppercase_snack_name ON snacks IS 'Prior to insert or update, force the name column to be all uppercase for data consistency';

