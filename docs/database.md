# Base Tables
---

## users
The full details of a beer league hockey team player

| Column     | Type                                            | Description                                                                                    |
| ---------- | ----------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique player identifier. Primary key for this table - assigned during insert to the table |
| first_name | TEXT NOT NULL                                   | The first name of the user                                                                     |
| last_name  | TEXT NOT NULL                                   | The last name of the user                                                                      |
| email      | TEXT NOT NULL                                   | The email address of the user                                                                  |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                        |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                   |


## teams
The full details of a beer league hockey team that a user might be a member of

| Column         | Type                                            | Description                                                                                  |
| -------------- | ----------------------------------------------- | -------------------------------------------------------------------------------------------- |
| id             | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique team identifier. Primary key for this table - assigned during insert to the table |
| name           | TEXT NOT NULL                                   | The name of the team                                                                         |
| rink           | rinks NOT NULL                                  | The primary rink that this team plays at - "rinks" is an enum type                           |
| level          | levels NOT NULL                                 | The level of the team - "levels" is an enum type                                             |
| primary_color  | TEXT                                            | The primary color of the team                                                                |
| seconary_color | TEXT                                            | The secondary color of the team                                                              |
| ternary_color  | TEXT                                            | The third color of the team                                                                  |
| logo_url       | TEXT                                            | The path to the team logo                                                                    |
| created_at     | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                      |
| updated_at     | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                 |


## allergies
Possible food related allergies

| Column     | Type                                            | Description                                                                                     |
| ---------- | ----------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique allergy identifier. Primary key for this table - assigned during insert to the table |
| name       | TEXT NOT NULL                                   | The basic name of the allergy                                                                   |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                         |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                    |


## snacks
A list of snacks with flavor profile, difficulty, and optional recipe URL

| Column     | Type                                            | Description                                                                                                    |
| ---------- | ----------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique snack identifier. Primary key for this table - assigned during insert to the table                  |
| name       | TEXT NOT NULL                                   | The name of the snack                                                                                          |
| sweet      | BOOLEAN NOT NULL                                | Whether the snack is considered sweet                                                                          |
| savory     | BOOLEAN NOT NULL                                | Whether the snack is considered savory                                                                         |
| difficulty | INT                                             | Arbitrary rating by Britni on the difficulty of the recipe - includes time to prepare and ingredients required |
| recipe_url | TEXT                                            | The url of the recipe                                                                                          |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                                        |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                                   |








# Linking Tables
---

## team_membership
Which users are members of which teams

| Column     | Type                                       | Description                                  |
| ---------- | ------------------------------------------ | -------------------------------------------- |
| team_id    | INT REFERENCES teams(id) ON DELETE CASCADE | The id from the teams table                  |
| user_id    | INT REFERENCES users(id) ON DELETE CASCADE | The id from the users table                  |
| created_at | TIMESTAMP DEFAULT now()                    | The time this row was created, UTC time      |
| updated_at | TIMESTAMP DEFAULT now()                    | The time this row was last updated, UTC time |

Note: PRIMARY KEY (team_id, user_id)


## user_allergies
Which users have allergies

| Column     | Type                                           | Description                                  |
| ---------- | ---------------------------------------------- | -------------------------------------------- |
| allergy_id | INT REFERENCES allergies(id) ON DELETE CASCADE | The id from the allergies table              |
| user_id    | INT REFERENCES users(id) ON DELETE CASCADE     | The id from the users table                  |
| created_at | TIMESTAMP DEFAULT now()                        | The time this row was created, UTC time      |
| updated_at | TIMESTAMP DEFAULT now()                        | The time this row was last updated, UTC time |

Note: PRIMARY KEY (allergy_id, user_id)


## snack_allergies
Which snacks have allergies

| Column     | Type                                           | Description                                  |
| ---------- | ---------------------------------------------- | -------------------------------------------- |
| allergy_id | INT REFERENCES allergies(id) ON DELETE CASCADE | The id from the allergies table              |
| snack_id   | INT REFERENCES snacks(id) ON DELETE CASCADE    | The id from the snacks table                 |
| created_at | TIMESTAMP DEFAULT now()                        | The time this row was created, UTC time      |
| updated_at | TIMESTAMP DEFAULT now()                        | The time this row was last updated, UTC time |

Note: PRIMARY KEY (allergy_id, snack_id)





# Logging Tables
---

## snack_log
A log of when snacks were made for teams

| Column     | Type                                            | Description                                                                                       |
| ---------- | ----------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique identifier of this snack offering. Primary key for this table - assigned during insert |
| snack_id   | INT REFERENCES snacks(id)                       | The id from the snacks table                                                                      |
| team_id    | INT REFERENCES teams(id)                        | The id from the teams table                                                                       |
| date_made  | DATE                                            | The date the snack was made                                                                       |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                           |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                      |





# Enums
---

- `rinks` - contains the names of rinks where the teams are located
- `levels` - contains the level indicators used by the rinks



# Functions
---

- `update_updated_at_column` - updates the value in the updated_at column when a row is updated. Used on all tables