
-- Create the audit log table
CREATE TABLE audit_log (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    table_name TEXT NOT NULL,
    action TEXT NOT NULL,
    changed_fields JSONB,
    db_user TEXT NOT NULL DEFAULT session_user,
    app_user TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
COMMENT ON TABLE audit_log IS 'The databse audit log. Tracks changes to columns using triggers in the database';
COMMENT ON COLUMN audit_log.id IS 'The unique identifier. Primary key for this table - assigned during insert to the table';
COMMENT ON COLUMN audit_log.table_name IS 'The name of the table where the change occurred, cannot be null';
COMMENT ON COLUMN audit_log.action IS 'Whether an INSERT, UPDATE or DELETE was performed, cannot be null';
COMMENT ON COLUMN audit_log.changed_fields IS 'For an UPDATE, shows the change in each field that was modified. For an INSERT, shows the valiues that were added. For a DELETE, shows the values that were removed.';
COMMENT ON COLUMN audit_log.db_user IS 'The database user that performed the change - set automatically';
COMMENT ON COLUMN audit_log.app_user IS 'The application user that performed the change - must be set by the application prior to the query. Defaults to "" if the app.current_user setting is not set';
COMMENT ON COLUMN audit_log.created_at IS 'The time this row was created, UTC time';
COMMENT ON COLUMN audit_log.updated_at IS 'The time this row was last updated, UTC time';

-- add the updated_at trigger to audit_log (the updated_at should never be used on this table, but let's be safe)
CREATE TRIGGER update_audit_log_updated_at
BEFORE UPDATE ON audit_log FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT ON TRIGGER update_audit_log_updated_at ON audit_log IS 'Update the updated_at timestamp column for the audit_log table';


-- Create the audit function
CREATE OR REPLACE FUNCTION audit_function()
RETURNS TRIGGER AS $$
DECLARE
    old_json JSONB;
    new_json JSONB;
    diff_json JSONB := '{}'::jsonb;
    key_name TEXT;
    old_val TEXT;
    new_val TEXT;
BEGIN
    -- Handle INSERT operations
    IF (TG_OP = 'INSERT') THEN
        INSERT INTO audit_log (table_name, action, changed_fields, app_user)
        VALUES (TG_TABLE_NAME, TG_OP, to_jsonb(NEW), NULLIF(current_setting('app.current_user', true), ''));
        RETURN NEW;
        
    -- Handle DELETE operations
    ELSIF (TG_OP = 'DELETE') THEN
        INSERT INTO audit_log (table_name, action, changed_fields, app_user)
        VALUES (TG_TABLE_NAME, TG_OP, to_jsonb(OLD), NULLIF(current_setting('app.current_user', true), ''));
        RETURN OLD;
        
    -- Handle UPDATE operations (Column-Level Isolate)
    ELSIF (TG_OP = 'UPDATE') THEN
        old_json := to_jsonb(OLD);
        new_json := to_jsonb(NEW);

        -- Loop through each column in the NEW record
        FOR key_name, new_val IN SELECT * FROM jsonb_each_text(new_json) LOOP
            old_val := old_json ->> key_name;
            
            -- Detect if the field value changed (handles NULL safely)
            IF old_val IS DISTINCT FROM new_val THEN
                diff_json := diff_json || jsonb_build_object(
                    key_name, 
                    jsonb_build_object('old', old_val, 'new', new_val)
                );
            END IF;
        END LOOP;

        -- Only log if actual differences were caught
        IF diff_json != '{}'::jsonb THEN
            INSERT INTO audit_log (table_name, action, changed_fields, app_user)
            VALUES (TG_TABLE_NAME, TG_OP, diff_json, NULLIF(current_setting('app.current_user', true), ''));
        END IF;
        
        RETURN NEW;
    END IF;
END;
$$ language 'plpgsql';


-- Teams - Add the trigger
CREATE TRIGGER audit_teams_trigger
AFTER INSERT OR UPDATE OR DELETE ON teams FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_teams_trigger ON teams IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- users - Add the trigger
CREATE TRIGGER audit_users_trigger
AFTER INSERT OR UPDATE OR DELETE ON users FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_users_trigger ON users IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- ingredients - Add the trigger
CREATE TRIGGER audit_ingredients_trigger
AFTER INSERT OR UPDATE OR DELETE ON ingredients FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_ingredients_trigger ON ingredients IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- snacks - Add the trigger
CREATE TRIGGER audit_snacks_trigger
AFTER INSERT OR UPDATE OR DELETE ON snacks FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_snacks_trigger ON snacks IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- suggested_allergies - Add the trigger
CREATE TRIGGER audit_suggested_allergies_trigger
AFTER INSERT OR UPDATE OR DELETE ON suggested_allergies FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_suggested_allergies_trigger ON suggested_allergies IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- team_membership - Add the trigger
CREATE TRIGGER audit_team_membership_trigger
AFTER INSERT OR UPDATE OR DELETE ON team_membership FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_team_membership_trigger ON team_membership IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- user_allergies - Add the trigger
CREATE TRIGGER audit_user_allergies_trigger
AFTER INSERT OR UPDATE OR DELETE ON user_allergies FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_user_allergies_trigger ON user_allergies IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- snack_ingredients - Add the trigger
CREATE TRIGGER audit_snack_ingredients_trigger
AFTER INSERT OR UPDATE OR DELETE ON snack_ingredients FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_snack_ingredients_trigger ON snack_ingredients IS 'After inserts, updates or deletes record audit information in the audit_log table';


-- snack_log - Add the trigger
CREATE TRIGGER audit_snack_log_trigger
AFTER INSERT OR UPDATE OR DELETE ON snack_log FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_snack_log_trigger ON snack_log IS 'After inserts, updates or deletes record audit information in the audit_log table';

-- user_snack_rankings - Add the trigger
CREATE TRIGGER audit_user_snack_rankings_trigger
AFTER INSERT OR UPDATE OR DELETE ON user_snack_rankings FOR EACH ROW EXECUTE PROCEDURE audit_function();
COMMENT ON TRIGGER audit_user_snack_rankings_trigger ON user_snack_rankings IS 'After inserts, updates or deletes record audit information in the audit_log table';