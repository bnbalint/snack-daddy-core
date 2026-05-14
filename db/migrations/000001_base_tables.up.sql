
-- Create the initial tables


-- this is the enum to hold the rinks
-- REMINDER - you can add more values later, but you can't remove values
-- this could be in its own table if we anticipated it needing to change more frequently or needing more details
CREATE TYPE rinks_enum AS ENUM ('BAIREL');
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


/*
 * teams table
 *
 */
CREATE TABLE teams (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    rink rinks_enum NOT NULL,
    level levels_enum NOT NULL,
    primary_color TEXT,
    seconary_color TEXT,
    ternary_color TEXT,
    logo_url TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE teams IS 'The full details of a beer league hockey team that a user might be a member of';
COMMENT ON COLUMN teams.id IS 'The unique team identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN teams.name IS 'The name of the team';
COMMENT ON COLUMN teams.rink IS 'The primary rink that this team plays at';
COMMENT ON COLUMN teams.level IS 'The level of the team';
COMMENT ON COLUMN teams.primary_color IS 'The primary color of the team';
COMMENT ON COLUMN teams.seconary_color IS 'The secondary color of the team';
COMMENT ON COLUMN teams.ternary_color IS 'The third color of the team';
COMMENT ON COLUMN teams.logo_url IS 'The path to the team logo';
COMMENT ON COLUMN teams.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN teams.updated_at IS 'The time this row was last updated, UTC time';

-- add the trigger to teams
CREATE TRIGGER update_teams_updated_at
BEFORE UPDATE ON teams FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_teams_updated_at ON teams IS 'Update the updated_at timestamp column for the teams table';



/*
 * users table
 * Team membership is recorded in the team_membership table
 * Allergies are recorded in the user_allergies table
 */
CREATE TABLE users (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE users IS 'The full details of a beer league hockey team player';
COMMENT ON COLUMN users.id IS 'The unique player identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN users.first_name IS 'The first name of the user';
COMMENT ON COLUMN users.last_name IS 'The last name of the user';
COMMENT ON COLUMN users.email IS 'The email address of the user';
COMMENT ON COLUMN users.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN users.updated_at IS 'The time this row was last updated, UTC time';

-- add the trigger to users
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_users_updated_at ON users IS 'Update the updated_at timestamp column for the users table';






/*
 * allergies table
 * this is a table because we will likely need to add more values over time
 */
CREATE TABLE allergies (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE allergies IS 'Possible food related allergies';
COMMENT ON COLUMN allergies.id IS 'The unique allergy identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN allergies.name IS 'The basic name of the allergy';
COMMENT ON COLUMN allergies.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN allergies.updated_at IS 'The time this row was last updated, UTC time';

-- add the trigger to allergies
CREATE TRIGGER update_allergies_updated_at
BEFORE UPDATE ON allergies FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_allergies_updated_at ON allergies IS 'Update the updated_at timestamp column for the allergies table';




/*
 * snacks table
 * Allergies are recorded in the snack_allergies table
 */
CREATE TABLE snacks (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    sweet BOOLEAN NOT NULL,
    savory BOOLEAN NOT NULL,
    difficulty INT,
    recipe_url TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE snacks IS 'Snacks and their descriptive qualities';
COMMENT ON COLUMN snacks.id IS 'The unique snack identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN snacks.name IS 'The name of the snack';
COMMENT ON COLUMN snacks.sweet IS 'Whether the snack is considered sweet';
COMMENT ON COLUMN snacks.savory IS 'Whether the snack is considered savory';
COMMENT ON COLUMN snacks.recipe_url IS 'The url of the recipe';
COMMENT ON COLUMN snacks.difficulty IS 'Arbitrary rating by Britni on the difficulty of the recipe - includes time to prepare and ingredients required. Scale of 1 (easy) to 10 (monstrous)';
COMMENT ON COLUMN snacks.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN snacks.updated_at IS 'The time this row was last updated, UTC time';

-- add the trigger to snacks
CREATE TRIGGER update_snacks_updated_at
BEFORE UPDATE ON snacks FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_snacks_updated_at ON snacks IS 'Update the updated_at timestamp column for the snacks table';



