# Base Tables
---

## users
The basic details of a user. Generally a beer league hockey team player, but can support spouses and fans as well. Further details (team membership, allergies) are handled via linking tables

| Column     | Type                                            | Description                                                                                    |
|------------|-------------------------------------------------|------------------------------------------------------------------------------------------------|
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique player identifier. Primary key for this table - assigned during insert to the table |
| first_name | TEXT NOT NULL                                   | The first name of the user                                                                     |
| last_name  | TEXT NOT NULL                                   | The last name of the user                                                                      |
| email      | TEXT NOT NULL UNIQUE                            | The email address of the user, must be unique                                                  |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                        |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                   |


## teams
The full details of a beer league hockey team that a user might be a member of

| Column          | Type                                            | Description                                                                                  |
|-----------------|-------------------------------------------------|----------------------------------------------------------------------------------------------|
| id              | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique team identifier. Primary key for this table - assigned during insert to the table |
| name            | TEXT NOT NULL UNIQUE                            | The name of the team, cannot be null, must be unique, converted to all CAPS upon saving      |
| rink            | rinks_enum NOT NULL                             | The primary rink that this team plays at - "rinks_enum" is an enum type                      |
| level           | levels_enum NOT NULL                            | The level of the team - "levels_enum" is an enum type                                        |
| primary_color   | TEXT NOT NULL                                   | The primary color of the team, saved as the hex color code, cannot be null                   |
| secondary_color | TEXT NOT NULL                                   | The secondary color of the team, saved as the hex color code, cannot be null                 |
| ternary_color   | TEXT NOT NULL DEFAULT ''                        | The third color of the team, saved as the hex color code, '' if no third color exists        |
| logo_url        | TEXT NOT NULL DEFAULT ''                        | The path to the team logo, '' if no image url exists                                         |
| created_at      | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                      |
| updated_at      | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                 |


## ingredients
Ingredients that a snack might have, or a user might be allergic to

| Column     | Type                                            | Description                                                                                         |
|------------|-------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique ingredient identifier. Primary key for this table - assigned during insert to the table  |
| name       | TEXT NOT NULL UNIQUE                            | The basic name of the ingredient, cannot be null, must be unique, converted to all CAPS upon saving |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                             |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                        |


## snacks
A list of snacks with flavor profile, difficulty, and optional recipe URL

| Column     | Type                                            | Description                                                                                                    |
|------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------|
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique snack identifier. Primary key for this table - assigned during insert to the table                  |
| name       | TEXT NOT NULL                                   | The name of the snack, cannot be null, must be unique, converted to all CAPS upon saving                       |
| sweet      | BOOLEAN NOT NULL                                | Whether the snack is considered sweet, true/false, cannot be null                                              |
| savory     | BOOLEAN NOT NULL                                | Whether the snack is considered savory, true/false, cannot be null                                             |
| difficulty | INT NOT NULL                                    | Arbitrary rating by Britni on the difficulty of the recipe - includes time to prepare and ingredients required |
| recipe_url | TEXT NOT NULL DEFAULT ''                        | The url of the recipe, '' if no url exists                                                                     |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                                        |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                                   |








# Linking Tables
---

## team_membership
Which users are members of which teams

| Column     | Type                                       | Description                                  |
|------------|--------------------------------------------|----------------------------------------------|
| team_id    | INT REFERENCES teams(id) ON DELETE CASCADE | The id from the teams table                  |
| user_id    | INT REFERENCES users(id) ON DELETE CASCADE | The id from the users table                  |
| created_at | TIMESTAMP DEFAULT now()                    | The time this row was created, UTC time      |
| updated_at | TIMESTAMP DEFAULT now()                    | The time this row was last updated, UTC time |

Note: PRIMARY KEY (team_id, user_id)


## user_allergies
Which users have allergies

| Column        | Type                                             | Description                                  |
|---------------|--------------------------------------------------|----------------------------------------------|
| ingredient_id | INT REFERENCES ingredients(id) ON DELETE CASCADE | The id from the ingredients table            |
| user_id       | INT REFERENCES users(id) ON DELETE CASCADE       | The id from the users table                  |
| created_at    | TIMESTAMP DEFAULT now()                          | The time this row was created, UTC time      |
| updated_at    | TIMESTAMP DEFAULT now()                          | The time this row was last updated, UTC time |

Note: PRIMARY KEY (ingredient_id, user_id)


## snack_ingredients
The ingredients that make up a snack

| Column        | Type                                             | Description                                  |
|---------------|--------------------------------------------------|----------------------------------------------|
| ingredient_id | INT REFERENCES ingredients(id) ON DELETE CASCADE | The id from the ingredients table            |
| snack_id      | INT REFERENCES snacks(id) ON DELETE CASCADE      | The id from the snacks table                 |
| created_at    | TIMESTAMP DEFAULT now()                          | The time this row was created, UTC time      |
| updated_at    | TIMESTAMP DEFAULT now()                          | The time this row was last updated, UTC time |

Note: PRIMARY KEY (ingredient_id, snack_id)





# Logging Tables
---

## snack_log
A log of when snacks were made for teams

| Column     | Type                                            | Description                                                                                       |
|------------|-------------------------------------------------|---------------------------------------------------------------------------------------------------|
| id         | BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY | The unique identifier of this snack offering. Primary key for this table - assigned during insert |
| snack_id   | INT REFERENCES snacks(id)                       | The id from the snacks table                                                                      |
| team_id    | INT REFERENCES teams(id)                        | The id from the teams table                                                                       |
| date_made  | DATE                                            | The date the snack was made                                                                       |
| created_at | TIMESTAMP DEFAULT now()                         | The time this row was created, UTC time                                                           |
| updated_at | TIMESTAMP DEFAULT now()                         | The time this row was last updated, UTC time                                                      |





# Enums
---

- `rinks_enum` - contains the names of rinks where the teams are located
- `levels_enum` - contains the level indicators used by the rinks



# Functions
---

- `update_updated_at_column` - updates the value in the updated_at column when a row is updated. Used on all tables
- `force_uppercase_name` - prior to INSERT or UPDATE, converts the value of the `name` column to all uppercase