# Admin

## Snack Entry
- Add a new snack to the snacks table
  - Name
  - Sweet (checkbox)
  - Savory (checkbox)
  - Recipe url (optional)
  - Difficulty (number, 1-10)
  - Ingredients (checkboxes)
- View existing snacks down the side (scroll bar)


## View / Modify Snack Difficulty
- ** Drag and drop into numbered buckets (1-10)  TODO **


## Team Entry
- Add a new team to the teams table
  - Team name
  - Rink
  - Level
  - Colors (color picker)
  - Upload image of logo
- View existing teams down the side (scroll bar)


## Ingredient Entry
- Add a new ingredient to the ingredients table
  - Name
- View existing ingredients down the side (scroll bar)


## Log Snack Delivery
- Log that a snack was made for a team
- Choose snack from dropdown
- Choose team from dropdown
- Choose date from datepicker


## View Snack Log
- View all snacks delivered
- Filter by team
- Filter by snack



# User Facing

## Register User
 - First name
 - Last name
 - Email
 - Team(s)  [or a spouse]
 - Allergies (checkboxes)
   - option for other with text --> this needs to be emailed to Britni so it can be added to the backend
 - Password


## Login User
 - ??

## Logout User
  - 

## Forgot Password
   -

## View / Modify User Settings
 - View user details
    - pulled from the users table (and linking tables)
    - first name, last name, email address
    - team membership
    - allergies
 - ** Modify allergies - TODO **
 - ** Modify team membership - TODO **


## View / Modify Snack Ranking
- Bucket for "Have not had"
- Bucket for "No Rating" (aka do not remember, but have had)
- Bucket for "My Top 3"
  - Limit to 3 in this bucket
  - Preserve order in the bucket
- Buckets for 1 (best) through 10 (worst) for scoring

