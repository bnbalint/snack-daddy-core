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

### Update snacks
- USAGE
  - Submit after modifying difficulty on [View / Modify Snack Difficulty](frontend.md#view--modify-snack-difficulty)
- PUT `/snacks`
- Body
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





## Snack Log
---

## GET all snack_log
- USAGE
  - View on [View Snack Log](frontend.md#view-snack-log)
- GET `/snack-log`
- Response
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


## Add to snack_log
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


## Rinks
---

## GET all rinks
- USAGE
  - Display on [Team Entry](./frontend.md#team-entry)
- GET `/rinks`
- Response
  ```json
  ["BAIREL","UPMC"]
  ```


## Levels
---

## GET all rinks
- USAGE
  - Display on [Team Entry](frontend.md#team-entry)
- GET `/levels`
- Response
  ```json
  ["D5","D4","D3"]
  ```







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