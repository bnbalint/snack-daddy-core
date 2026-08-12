# Endpoints

## Teams
---
### GET all teams
- USAGE:
  - Dropdown on [Log Snack Delivery](frontend.md#log-snack-delivery)
  - List on [Team Entry](frontend.md#team-entry)
- GET  `/teams`
- Response
  ```json
   [
     {
       "id": 1,
       "name": "Mules",
       "rink": "BAIREL",
       "level": "D5",
       "primary_color": "#b88907",
       "secondary_color": "#000000",
       "ternary_color": "#c42323",
       "logo_url": "",
       "created_at": "2026-07-01T00:00:01Z",
       "updated_at": "2026-07-01T00:00:01Z"
     }
   ]
  ```


### Add to teams
- USAGE
  - Submit on [View / Modify Snack Difficulty](frontend.md#view--modify-snack-difficulty)
- POST `/teams`
- Body
   ```json
     {
       "name": "Mules",
       "rink": "BAIREL",
       "level": "D5",
       "primary_color": "#b88907",
       "secondary_color": "#000000",
       "ternary_color": "#c42323",
       "logo_url": ""
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "name": "Mules",
      "rink": "BAIREL",
      "level": "D5",
      "primary_color": "#b88907",
      "secondary_color": "#000000",
      "ternary_color": "#c42323",
      "logo_url": "",
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required Team format
  - 400 → conflict because the team already exists






## Ingredients
---

### GET all ingredients
- USAGE
  - List on [Ingredient Entry](frontend.md#ingredient-entry)
- GET  `/ingredients`
- Response
   ```json
     [
       {
         "id": 1,
         "name": "Pecan",
         "created_at": "2026-07-01T00:00:01Z",
         "updated_at": "2026-07-01T00:00:01Z"
       },
       {
         "id": 2,
         "name": "Almond Milk",
         "created_at": "2026-07-01T00:00:01Z",
         "updated_at": "2026-07-01T00:00:01Z"
       }
     ]
   ```

### Add to ingredients
- USAGE
  - Submit on [Ingredient Entry](frontend.md#ingredient-entry)
- POST `\ingredients`
- Body
   ```json
     {
       "name": "Pecan"
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "name": "Pecan",
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required Ingredient format
  - 400 → conflict because the ingredient already exists (name must be unique)





## Users
---
### GET all users
- USAGE:
  - Future admin page of some kind (for viewing stats, registrations, etc)
- GET  `/users`
- Response
  ```json
    [
      {
        "id": 1,
        "first_name": "Roger",
        "last_name": "Hogwarts",
        "email": "r.h@gmail.com",
        "teams": [
          {
            "id": 1,
            "name": "Mules",
            "Rink": "BAIREL",
            "Level": "D5",
            "primary_color": "#b88907",
            "secondary_color": "#000000",
            "ternary_color": "#c42323",
            "logo_url": "",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
            }
        ],
        "allergies": [
          {
            "id": 1,
            "name": "Pecan",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      {
        "id": 2,
        "first_name": "Brandi",
        "last_name": "Hogwarts",
        "email": "r.h@gmail.com",
        "teams": [],
        "allergies": [],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
    ]
  ```

### GET user by id
- USAGE:
  - User-facing page for editing user details
- GET  `/users/{userId}`
- Response
  ```json
      {
        "id": 1,
        "first_name": "Roger",
        "last_name": "Hogwarts",
        "email": "r.h@gmail.com",
        "teams": [
          {
            "id": 1,
            "name": "Mules",
            "Rink": "BAIREL",
            "Level": "D5",
            "primary_color": "#b88907",
            "secondary_color": "#000000",
            "ternary_color": "#c42323",
            "logo_url": "",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
            }
        ],
        "allergies": [
          {
            "id": 1,
            "name": "Pecan",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
  ```
- Possible Errors
  - 400 → userId is null (Java) or unable to be converted to an integer (Go)
  - 400 → userId is less than or equal to 0
  - 404 → user was not found
  - 500 -> other error

### Add to users
- USAGE:
  - Create a new user
- POST `/users`
- Body
   ```json
     {
       "first_name": "Roger",
       "last_name": "Hogwarts",
       "email": "r.h@gmail.com",
       "teams": [
         {
           "id": 1,
           "name": "Mules",
           "Rink": "BAIREL",
           "Level": "D5",
           "primary_color": "#b88907",
           "secondary_color": "#000000",
           "ternary_color": "#c42323",
           "logo_url": "",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
           }
       ],
       "allergies": [
         {
           "id": 1,
           "name": "Pecan",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
         }
       ]
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "first_name": "Roger",
      "last_name": "Hogwarts",
      "email": "r.h@gmail.com",
      "teams": [
        {
          "id": 1,
          "name": "Mules",
          "Rink": "BAIREL",
          "Level": "D5",
          "primary_color": "#b88907",
          "secondary_color": "#000000",
          "ternary_color": "#c42323",
          "logo_url": "",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
          }
      ],
      "allergies": [
        {
          "id": 1,
          "name": "Pecan",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
        }
      ],
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required User format
  - 400 → conflict because the user already exists


### Update user
- USAGE:
  - Update a user (id and email cannot be modified)
- PUT `/users`
- Body
   ```json
     {
       "first_name": "Updated",
       "last_name": "Hogwarts",
       "email": "r.h@gmail.com",
       "teams": [
         {
           "id": 1,
           "name": "Mules",
           "Rink": "BAIREL",
           "Level": "D5",
           "primary_color": "#b88907",
           "secondary_color": "#000000",
           "ternary_color": "#c42323",
           "logo_url": "",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
           }
       ],
       "allergies": [
         {
           "id": 1,
           "name": "Pecan",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
         }
       ]
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "first_name": "Updated",
      "last_name": "Hogwarts",
      "email": "r.h@gmail.com",
      "teams": [
        {
          "id": 1,
          "name": "Mules",
          "Rink": "BAIREL",
          "Level": "D5",
          "primary_color": "#b88907",
          "secondary_color": "#000000",
          "ternary_color": "#c42323",
          "logo_url": "",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
          }
      ],
      "allergies": [
        {
          "id": 1,
          "name": "Pecan",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
        }
      ],
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-08-11T00:00:01Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required User format
  - 400 → database conflict


## Snacks
---

### GET all snacks
- USAGE:
  - Dropdown on [Log Snack Delivery](frontend.md#log-snack-delivery)
  - Display for [View / Modify Snack Difficulty](frontend.md#view--modify-snack-difficulty)
  - List on [Snack Entry](frontend.md#snack-entry)
- GET  `/snacks`
- Response
  ```json
    [
      {
        "id": 1,
        "name": "Rice Crispie Treat",
        "sweet": true,
        "savory": false,
        "difficulty": 2,
        "recipe_url": "",
        "ingredients": [
          {
            "id": 4,
            "name": "Rice Crispy Cereal",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 5,
            "name": "Margarine",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 6,
            "name": "Marshmallow",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
    ]
  ```


### Add to snacks
- USAGE
  - Submit on [Snack Entry](frontend.md#snack-entry)
- POST `/snacks`
- Body
   ```json
     {
       "name": "Rice Crispie Treat",
       "sweet": true,
       "savory": false,
       "difficulty": 2,
       "recipe_url": "",
       "ingredients": [
         {
           "id": 4,
           "name": "Rice Crispy Cereal",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
         },
         {
           "id": 5,
           "name": "Margarine",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
         },
         {
           "id": 6,
           "name": "Marshmallow",
           "created_at": "2026-07-01T00:00:01Z",
           "updated_at": "2026-07-01T00:00:01Z"
         }
       ]
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "name": "Rice Crispie Treat",
      "sweet": true,
      "savory": false,
      "difficulty": 2,
      "recipe_url": "",
      "ingredients": [
        {
          "id": 4,
          "name": "Rice Crispy Cereal",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
        },
        {
          "id": 5,
          "name": "Margarine",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
        },
        {
          "id": 6,
          "name": "Marshmallow",
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
        }
      ],
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required Snack format
  - 400 → conflict because the snack already exists


### Update snacks
- USAGE
  - Submit after modifying difficulty on [View / Modify Snack Difficulty](frontend.md#view--modify-snack-difficulty)
- PUT `/snacks`
- Body = list of Snacks
   ```json
     [
       {
         "id": 1,
         "name": "Rice Crispie Treat",
         "sweet": true,
         "savory": false,
         "difficulty": 5,
         "recipe_url": "",
         "ingredients": [
           {
             "id": 4,
             "name": "Rice Crispy Cereal",
             "created_at": "2026-07-01T00:00:01Z",
             "updated_at": "2026-07-01T00:00:01Z"
           },
           {
             "id": 5,
             "name": "Margarine",
             "created_at": "2026-07-01T00:00:01Z",
             "updated_at": "2026-07-01T00:00:01Z"
           },
           {
             "id": 6,
             "name": "Marshmallow",
             "created_at": "2026-07-01T00:00:01Z",
             "updated_at": "2026-07-01T00:00:01Z"
           }
         ],
         "created_at": "2026-07-01T00:00:01Z",
         "updated_at": "2026-07-01T00:00:01Z"
       }
     ]
   ```
- Response = list of Snacks
  ```json
    [
      {
        "id": 1,
        "name": "Rice Crispie Treat",
        "sweet": true,
        "savory": false,
        "difficulty": 2,
        "recipe_url": "",
        "ingredients": [
          {
            "id": 4,
            "name": "Rice Crispy Cereal",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 5,
            "name": "Margarine",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 6,
            "name": "Marshmallow",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
    ]
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required Team format


  
## Suggested Allergies
---

### GET all suggested_allergies
- USAGE
  - Admin page to [View Suggested Allergies](./frontend.md#view-suggested-allergies)
- GET  `/suggested-allergies`
- Response
   ```json
     [
       {
         "id": 1,
         "name": "Pine nut",
         "created_at": "2026-07-01T00:00:01Z",
         "updated_at": "2026-07-01T00:00:01Z"
       },
       {
         "id": 2,
         "name": "Gluten",
         "created_at": "2026-07-01T00:00:01Z",
         "updated_at": "2026-07-01T00:00:01Z"
       }
     ]
   ```

### Add to suggested_allergies
- USAGE
  - Submit on [Register User](frontend.md#register-user) for submitting an allergy that is not listed
- POST `\suggested-allergies`
- Body
   ```json
     {
       "name": "Pine nut"
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "name": "Pine nut",
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required SuggestedAllergy format



## Snack Log
---

### GET all snack_log
- USAGE
  - View on [View Snack Log](frontend.md#view-snack-log) Admin page
- GET `/snack-log`
- Response = list of SnackLog entries
  ```json
   [
    {
      "id": 1,
      "snack_id": 1,
      "snack": {
        "id": 1,
        "name": "Rice Crispie Treat",
        "sweet": true,
        "savory": false,
        "difficulty": 2,
        "recipe_url": "",
        "ingredients": [
          {
            "id": 4,
            "name": "Rice Crispy Cereal",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 5,
            "name": "Margarine",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 6,
            "name": "Marshmallow",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 7,
            "name": "Vanilla",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      "team_id": 1,
      "team": {
        "id": 1,
        "name": "Mules",
        "rink": "BAIREL",
        "level": "D5",
        "primary_color": "#b88907",
        "secondary_color": "#000000",
        "ternary_color": "#c42323",
        "logo_url": "logo.com",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      "date_made": "2026-07-20",
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    },
    {
      "id": 2,
      "snack_id": 1,
      "snack": {
        "id": 1,
        "name": "Rice Crispie Treat",
        "sweet": true,
        "savory": false,
        "difficulty": 2,
        "recipe_url": "",
        "ingredients": [
          {
            "id": 4,
            "name": "Rice Crispy Cereal",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 5,
            "name": "Margarine",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 6,
            "name": "Marshmallow",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 7,
            "name": "Vanilla",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      "team_id": 1,
      "team": {
        "id": 2,
        "name": "Monsters",
        "rink": "BAIREL",
        "level": "D4",
        "primary_color": "#2c54c0",
        "secondary_color": "#000000",
        "ternary_color": "#e967b7",
        "logo_url": "logo.com",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      "date_made": "2026-07-20",
      "created_at": "2026-07-01T00:00:01Z",
      "updated_at": "2026-07-01T00:00:01Z"
    }
   ]
  ```


### Add to snack_log
- USAGE
  - Submit on [Log Snack Delivery](frontend.md#log-snack-delivery)
- POST `/snack-log`
- Body
   ```json
     {
       "snack_id": 1,
       "team_id": 1,
       "date_made": "2026-06-25"
     }
   ```
- Response
  ```json
    {
      "id": 1,
      "snack_id": 1,
      "snack": {
        "id": 1,
        "name": "Rice Crispie Treat",
        "sweet": true,
        "savory": false,
        "difficulty": 2,
        "recipe_url": "",
        "ingredients": [
          {
            "id": 4,
            "name": "Rice Crispy Cereal",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 5,
            "name": "Margarine",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 6,
            "name": "Marshmallow",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          },
          {
            "id": 7,
            "name": "Vanilla",
            "created_at": "2026-07-01T00:00:01Z",
            "updated_at": "2026-07-01T00:00:01Z"
          }
        ],
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      "team_id": 1,
      "team": {
        "id": 1,
        "name": "Mules",
        "rink": "BAIREL",
        "level": "D5",
        "primary_color": "#b88907",
        "secondary_color": "#000000",
        "ternary_color": "#c42323",
        "logo_url": "logo.com",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      },
      "date_made": "2026-07-20",
      "created_at": "2026-07-25T12:00:00Z",
      "updated_at": "2026-07-25T12:00:00Z"
    }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required SnackLog format





## Rinks
---

### GET all rinks
- USAGE
  - Display on [Team Entry](./frontend.md#team-entry)
- GET `/rinks`
- Response
  ```json
  ["BAIREL","UPMC"]
  ```




## Levels
---

### GET all levels
- USAGE
  - Display on [Team Entry](frontend.md#team-entry)
- GET `/levels`
- Response
  ```json
  ["D5","D4","D3"]
  ```






## Snack Rankings
---
### GET all snack rankings (all users)
- USAGE:
  - Future admin page of some kind (for viewing all data)
- GET  `/snack-rankings`
- Response
  ```json
    [
      {
        "snack_id": 1,
        "snack": {
          "id": 1,
          "name": "Rice Crispie Treat",
          "sweet": true,
          "savory": false,
          "difficulty": 2,
          "recipe_url": "",
          "ingredients": [
            {
              "id": 4,
              "name": "Rice Crispy Cereal",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 5,
              "name": "Margarine",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 6,
              "name": "Marshmallow",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 7,
              "name": "Vanilla",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            }
          ],
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
      },
        "user_id": 1,
        "user": {
            "first_name": "Roger",
            "last_name": "Hogwarts",
            "email": "r.h@gmail.com",
            "teams": [
              {
                "id": 1,
                "name": "Mules",
                "Rink": "BAIREL",
                "Level": "D5",
                "primary_color": "#b88907",
                "secondary_color": "#000000",
                "ternary_color": "#c42323",
                "logo_url": "",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
                }
            ],
            "allergies": [
              {
                "id": 1,
                "name": "Pecan",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
              }
            ]
          },
        "rank": "RANK_10",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
    ]
  ```

### GET all snack rankings (one user)
- USAGE:
  - As a user, view snacks and their rankings on the drag-and-drop ranking page
- GET  `/users/{userId}/snack-rankings`
- Response
  ```json
    [
      {
        "snack_id": 1,
        "snack": {
          "id": 1,
          "name": "Rice Crispie Treat",
          "sweet": true,
          "savory": false,
          "difficulty": 2,
          "recipe_url": "",
          "ingredients": [
            {
              "id": 4,
              "name": "Rice Crispy Cereal",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 5,
              "name": "Margarine",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 6,
              "name": "Marshmallow",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 7,
              "name": "Vanilla",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            }
          ],
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
      },
        "user_id": 1,
        "user": {
            "first_name": "Roger",
            "last_name": "Hogwarts",
            "email": "r.h@gmail.com",
            "teams": [
              {
                "id": 1,
                "name": "Mules",
                "Rink": "BAIREL",
                "Level": "D5",
                "primary_color": "#b88907",
                "secondary_color": "#000000",
                "ternary_color": "#c42323",
                "logo_url": "",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
                }
            ],
            "allergies": [
              {
                "id": 1,
                "name": "Pecan",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
              }
            ]
          },
        "rank": "RANK_10",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
    ]
  ```

### Add to snack ranking
- USAGE:
  - Create a new snack ranking
- POST `/snack-rankings`
- Body
   ```json
      {
        "snack_id": 1,
        "snack": {
          "id": 1,
          "name": "Rice Crispie Treat",
          "sweet": true,
          "savory": false,
          "difficulty": 2,
          "recipe_url": "",
          "ingredients": [
            {
              "id": 4,
              "name": "Rice Crispy Cereal",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 5,
              "name": "Margarine",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 6,
              "name": "Marshmallow",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 7,
              "name": "Vanilla",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            }
          ],
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
      },
        "user_id": 1,
        "user": {
            "first_name": "Roger",
            "last_name": "Hogwarts",
            "email": "r.h@gmail.com",
            "teams": [
              {
                "id": 1,
                "name": "Mules",
                "Rink": "BAIREL",
                "Level": "D5",
                "primary_color": "#b88907",
                "secondary_color": "#000000",
                "ternary_color": "#c42323",
                "logo_url": "",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
                }
            ],
            "allergies": [
              {
                "id": 1,
                "name": "Pecan",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
              }
            ]
          },
        "rank": "RANK_10",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
   ```
- Response
  ```json
      {
        "snack_id": 1,
        "snack": {
          "id": 1,
          "name": "Rice Crispie Treat",
          "sweet": true,
          "savory": false,
          "difficulty": 2,
          "recipe_url": "",
          "ingredients": [
            {
              "id": 4,
              "name": "Rice Crispy Cereal",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 5,
              "name": "Margarine",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 6,
              "name": "Marshmallow",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 7,
              "name": "Vanilla",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            }
          ],
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
      },
        "user_id": 1,
        "user": {
            "first_name": "Roger",
            "last_name": "Hogwarts",
            "email": "r.h@gmail.com",
            "teams": [
              {
                "id": 1,
                "name": "Mules",
                "Rink": "BAIREL",
                "Level": "D5",
                "primary_color": "#b88907",
                "secondary_color": "#000000",
                "ternary_color": "#c42323",
                "logo_url": "",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
                }
            ],
            "allergies": [
              {
                "id": 1,
                "name": "Pecan",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
              }
            ]
          },
        "rank": "RANK_10",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
  ```
- Possible Errors
  - 415 → JSON body was unable to be parsed into the required UserSnackRanking format
  - 400 → conflict because the user/snack entry already exists (would need to use a PUT to update this entry)



### Update UserSnackRankings
- USAGE:
  - Save on the drag-and-drop page for updating Snack Rankings
- PUT `/snack-rankings`
- Body
   ```json
      [
        {"snack_id": 1,
        "snack": {
          "id": 1,
          "name": "Rice Crispie Treat",
          "sweet": true,
          "savory": false,
          "difficulty": 2,
          "recipe_url": "",
          "ingredients": [
            {
              "id": 4,
              "name": "Rice Crispy Cereal",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 5,
              "name": "Margarine",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 6,
              "name": "Marshmallow",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 7,
              "name": "Vanilla",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            }
          ],
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
      },
        "user_id": 1,
        "user": {
            "first_name": "Roger",
            "last_name": "Hogwarts",
            "email": "r.h@gmail.com",
            "teams": [
              {
                "id": 1,
                "name": "Mules",
                "Rink": "BAIREL",
                "Level": "D5",
                "primary_color": "#b88907",
                "secondary_color": "#000000",
                "ternary_color": "#c42323",
                "logo_url": "",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
                }
            ],
            "allergies": [
              {
                "id": 1,
                "name": "Pecan",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
              }
            ]
          },
        "rank": "RANK_10",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
      ]
   ```
- Response
  ```json
      [
        {
        "snack_id": 1,
        "snack": {
          "id": 1,
          "name": "Rice Crispie Treat",
          "sweet": true,
          "savory": false,
          "difficulty": 2,
          "recipe_url": "",
          "ingredients": [
            {
              "id": 4,
              "name": "Rice Crispy Cereal",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 5,
              "name": "Margarine",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 6,
              "name": "Marshmallow",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            },
            {
              "id": 7,
              "name": "Vanilla",
              "created_at": "2026-07-01T00:00:01Z",
              "updated_at": "2026-07-01T00:00:01Z"
            }
          ],
          "created_at": "2026-07-01T00:00:01Z",
          "updated_at": "2026-07-01T00:00:01Z"
      },
        "user_id": 1,
        "user": {
            "first_name": "Roger",
            "last_name": "Hogwarts",
            "email": "r.h@gmail.com",
            "teams": [
              {
                "id": 1,
                "name": "Mules",
                "Rink": "BAIREL",
                "Level": "D5",
                "primary_color": "#b88907",
                "secondary_color": "#000000",
                "ternary_color": "#c42323",
                "logo_url": "",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
                }
            ],
            "allergies": [
              {
                "id": 1,
                "name": "Pecan",
                "created_at": "2026-07-01T00:00:01Z",
                "updated_at": "2026-07-01T00:00:01Z"
              }
            ]
          },
        "rank": "RANK_10",
        "created_at": "2026-07-01T00:00:01Z",
        "updated_at": "2026-07-01T00:00:01Z"
      }
    ]
  ```
- Possible Errors
  - 415 -> JSON body was unable to be parsed into the required UserSnackRanking format
  - 400 -> conflict during the database transaction





## Other
---

## Upload logo
- USAGE
  - Upload button on [Team Entry](frontend.md#team-entry)
- POST `/upload-image`
- Body
   ```json
   ```
- Response
    ```json
    ```