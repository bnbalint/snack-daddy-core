
-- Create the linking tables




/*
 * team_membership table
 *
 */
CREATE TABLE team_membership (
    team_id INT REFERENCES teams(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (team_id, user_id)
);
COMMENT ON TABLE team_membership IS 'Which users are members of which teams';
COMMENT ON COLUMN team_membership.team_id IS 'The id from the teams table';
COMMENT ON COLUMN team_membership.user_id IS 'The id from the users table';
COMMENT ON COLUMN team_membership.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN team_membership.updated_at IS 'The time this row was last updated, UTC time';


-- add the trigger to team_membership
CREATE TRIGGER update_team_membership_updated_at
BEFORE UPDATE ON team_membership FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_team_membership_updated_at ON team_membership IS 'Update the updated_at timestamp column for the team_membership table';




/*
 * user_allergies table
 *
 */
CREATE TABLE user_allergies (
    allergy_id INT REFERENCES allergies(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (allergy_id, user_id)
);
COMMENT ON TABLE user_allergies IS 'Which users have allergies';
COMMENT ON COLUMN user_allergies.allergy_id IS 'The id from the allergies table';
COMMENT ON COLUMN user_allergies.user_id IS 'The id from the users table';
COMMENT ON COLUMN user_allergies.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN user_allergies.updated_at IS 'The time this row was last updated, UTC time';


-- add the trigger to user_allergies
CREATE TRIGGER update_user_allergies_updated_at
BEFORE UPDATE ON user_allergies FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_user_allergies_updated_at ON user_allergies IS 'Update the updated_at timestamp column for the user_allergies table';




/*
 * snack_allergies table
 *
 */
CREATE TABLE snack_allergies (
    allergy_id INT REFERENCES allergies(id) ON DELETE CASCADE,
    snack_id INT REFERENCES snacks(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (allergy_id, snack_id)
);
COMMENT ON TABLE snack_allergies IS 'Which snacks have allergies';
COMMENT ON COLUMN snack_allergies.allergy_id IS 'The id from the allergies table';
COMMENT ON COLUMN snack_allergies.snack_id IS 'The id from the snack table';
COMMENT ON COLUMN snack_allergies.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN snack_allergies.updated_at IS 'The time this row was last updated, UTC time';


-- add the trigger to snack_allergies
CREATE TRIGGER update_snack_allergies_updated_at
BEFORE UPDATE ON snack_allergies FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_snack_allergies_updated_at ON snack_allergies IS 'Update the updated_at timestamp column for the snack_allergies table';



/*
 * snack_log
 *
 */
CREATE TABLE snack_log (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    snack_id INT REFERENCES snacks(id),
    team_id INT REFERENCES teams(id),
    date_made DATE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE snack_log IS 'A log of when snacks were made for teams';
COMMENT ON COLUMN snack_log.id IS 'The unique identifier of this snack offering. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN snack_log.snack_id IS 'The id from the snacks table';
COMMENT ON COLUMN snack_log.team_id IS 'The id from the teams table';
COMMENT ON COLUMN snack_log.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN snack_log.updated_at IS 'The time this row was last updated, UTC time';


-- add the trigger to snack_log
CREATE TRIGGER update_snack_log_updated_at
BEFORE UPDATE ON snack_log FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_snack_log_updated_at ON snack_log IS 'Update the updated_at timestamp column for the snack_log table';