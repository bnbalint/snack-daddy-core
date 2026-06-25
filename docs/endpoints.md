# Endpoints

## Teams
---
### GET all teams
 - USAGE:
   - Dropdown on [Log Snack Delivery](./frontend.md#log-snack-delivery)
   - USAGE: List on [Team Entry](./frontend.md#team-entry)
 - GET  `/teams`
 - Response
   ```json
    [
      {
        "ID": 1,
        "Name": "Mules",
        "Rink": "BAIREL",
        "Level": "D5",
        "PrimaryColor": "#b88907",
        "SecondaryColor": "#000000",
        "TernaryColor": "#c42323",
        "LogoUrl": "",
        "Created_At": "2026-06-25T12:00:00Z",
        "Updated_At": "2026-06-25T12:00:00Z"
      }
    ]
   ```


### Add to teams
 - USAGE: Submit on [View / Modify Snack Difficulty](./frontend.md#view--modify-snack-difficulty)
 - POST `/teams`
 - Body
    ```json
      {
        "Name": "Mules",
        "Rink": "BAIREL",
        "Level": "D5",
        "PrimaryColor": "#b88907",
        "SecondaryColor": "#000000",
        "TernaryColor": "#c42323",
        "LogoUrl": ""
        }
    ```
- Response
  ```json
    {
      "ID": 1,
      "Name": "Mules",
      "Rink": "BAIREL",
      "Level": "D5",
      "PrimaryColor": "#b88907",
      "SecondaryColor": "#000000",
      "TernaryColor": "#c42323",
      "LogoUrl": "",
      "Created_At": "2026-06-25T12:00:00Z",
      "Updated_At": "2026-06-25T12:00:00Z"
      }
  ```


## Ingredients
---

### GET all ingredients
 - USAGE: List on [Ingredient Entry](./frontend.md#ingredient-entry)
 - GET  `/ingredients`
 - Response
    ```json
      [
        {
          "ID": 1,
          "Name": "Pecan",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        },
        {
          "ID": 2,
          "Name": "Almond Milk",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        }
      ]
    ```

### Add to ingredients
 - USAGE: Submit on [Ingredient Entry](./frontend.md#ingredient-entry)
 - POST `\ingredients`
 - Body
    ```json
      {
        "Name": "Pecan"
      }
    ```
- Response
  ```json
    {
      "ID": 1,
      "Name": "Pecan",
      "Created_At": "2026-06-25T12:00:00Z",
      "Updated_At": "2026-06-25T12:00:00Z"
    }
  ```


## Users
---
### GET all users
 - USAGE:
 - GET  `/users`
- Response
  ```json
    [
      {
        "ID": 1,
        "First_Name": "Roger",
        "Last_Name": "Hogwarts",
        "Email": "r.h@gmail.com",
        "Teams": [
          {
            "ID": 1,
            "Name": "Mules",
            "Rink": "BAIREL",
            "Level": "D5",
            "PrimaryColor": "#b88907",
            "SecondaryColor": "#000000",
            "TernaryColor": "#c42323",
            "LogoUrl": "",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
            }
        ],
        "Allergies": [
          {
            "ID": 1,
            "Name": "Pecan",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          }
        ],
        "Created_At": "2026-06-25T12:00:00Z",
        "Updated_At": "2026-06-25T12:00:00Z"
      },
      {
        "ID": 2,
        "First_Name": "Brandi",
        "Last_Name": "Hogwarts",
        "Email": "r.h@gmail.com",
        "Teams": [],
        "Allergies": [],
        "Created_At": "2026-06-25T12:00:00Z",
        "Updated_At": "2026-06-25T12:00:00Z"
      },
    ]
  ```


### Add to users
 - USAGE: 
 - POST `/users`
 - Body
    ```json
      {
        "First_Name": "Roger",
        "Last_Name": "Hogwarts",
        "Email": "r.h@gmail.com",
        "Teams": [
          {
            "ID": 1,
            "Name": "Mules",
            "Rink": "BAIREL",
            "Level": "D5",
            "PrimaryColor": "#b88907",
            "SecondaryColor": "#000000",
            "TernaryColor": "#c42323",
            "LogoUrl": "",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
            }
        ],
        "Allergies": [
          {
            "ID": 1,
            "Name": "Pecan",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          }
        ]
      }
    ```
- Response
  ```json
    {
      "ID": 1,
      "First_Name": "Roger",
      "Last_Name": "Hogwarts",
      "Email": "r.h@gmail.com",
      "Teams": [
        {
          "ID": 1,
          "Name": "Mules",
          "Rink": "BAIREL",
          "Level": "D5",
          "PrimaryColor": "#b88907",
          "SecondaryColor": "#000000",
          "TernaryColor": "#c42323",
          "LogoUrl": "",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
          }
      ],
      "Allergies": [
        {
          "ID": 1,
          "Name": "Pecan",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        }
      ],
      "Created_At": "2026-06-25T12:00:00Z",
      "Updated_At": "2026-06-25T12:00:00Z"
    }
  ```



## Snacks
---

### GET all snacks
 - USAGE:
   - Dropdown on [Log Snack Delivery](./frontend.md#log-snack-delivery)
   - Display for [View / Modify Snack Difficulty](./frontend.md#view--modify-snack-difficulty)
   - List on [Snack Entry](./frontend.md#snack-entry)
 - GET  `/snacks`
- Response
  ```json
    [
      {
        "ID": 1,
        "Name": "Rice Crispie Treat",
        "Sweet": true,
        "Savory": false,
        "Difficulty": 2,
        "Recipe_Url": "",
        "Ingredients": [
          {
            "ID": 4,
            "Name": "Rice Crispy Cereal",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
          {
            "ID": 5,
            "Name": "Margarine",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
          {
            "ID": 6,
            "Name": "Marshmallow",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
        ],
        "Created_At": "2026-06-25T12:00:00Z",
        "Updated_At": "2026-06-25T12:00:00Z"
      }
    ]
  ```


### Add to snacks
 - USAGE: Submit on [Snack Entry](./frontend.md#snack-entry)
 - POST `/snacks`
 - Body
    ```json
      {
        "Name": "Rice Crispie Treat",
        "Sweet": true,
        "Savory": false,
        "Difficulty": 2,
        "Recipe_Url": "",
        "Ingredients": [
          {
            "ID": 4,
            "Name": "Rice Crispy Cereal",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
          {
            "ID": 5,
            "Name": "Margarine",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
          {
            "ID": 6,
            "Name": "Marshmallow",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
        ]
      }
    ```
- Response
  ```json
    {
      "ID": 1,
      "Name": "Rice Crispie Treat",
      "Sweet": true,
      "Savory": false,
      "Difficulty": 2,
      "Recipe_Url": "",
      "Ingredients": [
        {
          "ID": 4,
          "Name": "Rice Crispy Cereal",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        },
        {
          "ID": 5,
          "Name": "Margarine",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        },
        {
          "ID": 6,
          "Name": "Marshmallow",
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        },
      ],
      "Created_At": "2026-06-25T12:00:00Z",
      "Updated_At": "2026-06-25T12:00:00Z"
      }
  ```

### Update snacks
 - USAGE: Submit after modifying difficulty on [View / Modify Snack Difficulty](./frontend.md#view--modify-snack-difficulty)
 - PUT `/snacks`
 - Body
    ```json
      [
        {
          "ID": 1,
          "Name": "Rice Crispie Treat",
          "Sweet": true,
          "Savory": false,
          "Difficulty": 5,
          "Recipe_Url": "",
          "Ingredients": [
            {
              "ID": 4,
              "Name": "Rice Crispy Cereal",
              "Created_At": "2026-06-25T12:00:00Z",
              "Updated_At": "2026-06-25T12:00:00Z"
            },
            {
              "ID": 5,
              "Name": "Margarine",
              "Created_At": "2026-06-25T12:00:00Z",
              "Updated_At": "2026-06-25T12:00:00Z"
            },
            {
              "ID": 6,
              "Name": "Marshmallow",
              "Created_At": "2026-06-25T12:00:00Z",
              "Updated_At": "2026-06-25T12:00:00Z"
            },
          ],
          "Created_At": "2026-06-25T12:00:00Z",
          "Updated_At": "2026-06-25T12:00:00Z"
        }
      ]
    ```
- Response
  ```json
    [
      {
        "ID": 1,
        "Name": "Rice Crispie Treat",
        "Sweet": true,
        "Savory": false,
        "Difficulty": 5,
        "Recipe_Url": "",
        "Ingredients": [
          {
            "ID": 4,
            "Name": "Rice Crispy Cereal",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
          {
            "ID": 5,
            "Name": "Margarine",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
          {
            "ID": 6,
            "Name": "Marshmallow",
            "Created_At": "2026-06-25T12:00:00Z",
            "Updated_At": "2026-06-25T12:00:00Z"
          },
        ],
        "Created_At": "2026-06-25T12:00:00Z",
        "Updated_At": "2026-06-25T12:00:00Z"
      }
    ]
  ```





## Snack Log
---

## GET all snack_log
 - USAGE: View on [View Snack Log](frontend.md#view-snack-log)
 - GET `/snack-log`
- Response
  ```json
    [
      {

      }
    ]
  ```


## Add to snack_log
 - USAGE: Submit on [Log Snack Delivery](./frontend.md#log-snack-delivery)
 - POST `/snack-log`
 - Body
    ```json
      {
        "Snack_ID": 1,
        "Team_ID": 1,
        "Date_Made": "2026-06-25T12:00:00Z"
      }
    ```
- Response
  ```json
    {
      "ID": 1,
      "Snack_ID": 1,
      "Team_ID": 1,
      "Date_Made": "2026-06-20T12:00:00Z",
      "Created_At": "2026-06-25T12:00:00Z",
      "Updated_At": "2026-06-25T12:00:00Z"
    }
  ```





## Other
---

## Upload logo
 - USAGE: Upload button on [Team Entry](./frontend.md#team-entry)
 - POST `/upload-image`
 - Body
    ```
    ```
- Response
    ```
    ```