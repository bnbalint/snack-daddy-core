# Overview

`snack-daddy-core` serves as the backend core for the SnackDaddy project. <br>
This project is being slowly written in Golang as a means to learn Golang.

# What is SnackDaddy?

Weird name, right? That's what my husband's hockey team has nicknamed him because I always send baked goods for the team for after the game.

SnackDaddy, the code, will be a microservice to support a small frontend application for the hockey players to rate the baked goods. <br>
Because every time he comes home from a game, he swears three people claimed the latest snack is their new favorite. <br>
This seems suspicious to me - they can't have that many favorites! <br>
If only we could allow players to rank the snacks and then allow me to see this data ....

# Basic Functionality

- Ingredients
  - Support the **admin** adding/editing of ingredients
- Teams
  - Support the **admin** adding/editing of teams
- Snacks
  - Support the **admin** adding/editing of snacks
  - Each snack should contain properties to allow for future sorting/analytics on ingredients, difficulty and allergies
- Users
  - Support the **user** creation/editing of their own user
  - Support allergy indication for users
- Snack Log
  - Support the **admin** adding/viewing of when snacks were made for specific teams
- Snack Ranking
  - Support the **user** ranking of snacks that have been logged as made for their team


## Database
Information on the database tables can be found in [docs/database.md](./docs/database.md)


## Endpoints
An overview of endpoints can be found in [docs/endpoints.md](./docs/endpoints.md)


## Frontend / User Design
An overview of the user design can be found in [docs/frontend.md](./docs/frontend.md). This documentation will provide context of how the overall product will work.


## Adding a new Rink or Level

Note that the options for Rinks and Levels are managed in enums **both** at the database level and the code level. <br>
To add a new Rink/Level, you would need to add a database migration, adding the value to the database enum and then also do a code release to update the enum in the code <br>.
Heavy handed? Maybe. But we would like to force data consistency, and nobody can spell BAIREL. <br>
Also, we don't *actually* expect these to ever change.

[Why both? - So that a manual manipulation of the database has forced consistency in addition to the application itself]


# Release Notes

## 1.0.0
- Date: 2026-06-23
- Changes
  - Still in development --> where does the version even go LOL