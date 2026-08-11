package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for user

// Get all users
// Returns:
//
//	200 and a list of all users
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllUsers(ctx echo.Context) error {
	server.Logger.Debug("GetAllUsers")

	users, err := server.DB.GetAllUsers(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All users = %v", users))
	return ctx.JSON(http.StatusOK, users)
}

// Get single user by ID
// Returns:
//
//	  400 if the userId is not a valid integer
//		 400 if the userId is 0 or negative
//		 404 if the userId does not exist
//		 500 for all other errors
//		 200 and the user
func (server *SnackDaddyEchoServer) GetUserById(ctx echo.Context) error {

	// get the id value from the path
	id := ctx.Param("userId")
	server.Logger.Debug(fmt.Sprintf("GetUserById - userId = %v", id))

	// convert to integer
	userId, err := strconv.Atoi(id)

	if err != nil {
		server.Logger.Info(fmt.Sprintf("UserId (%v) must be a valid integer", userId))
		return ctx.JSON(http.StatusBadRequest, "UserId must be a valid integer")
	}

	if userId <= 0 {
		server.Logger.Info(fmt.Sprintf("UserId (%v) must be greater than 0", userId))
		return ctx.JSON(http.StatusBadRequest, "UserId must be greater than 0")
	}

	user, err := server.DB.GetUserById(ctx.Request().Context(), userId)
	if err != nil {

		// return 404 if the user was not found
		var notFoundError = &database_errors.NotFoundError{}
		if errors.As(err, &notFoundError) {
			return ctx.JSON(http.StatusNotFound, err)
		}

		// return 500 for all other error
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("User = %v", user))
	return ctx.JSON(http.StatusOK, user)
}

// Add a new user
// Expects the user to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created user object
//	415 if the body cannot be correctly parsed into a user object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddUser(ctx echo.Context) error {
	server.Logger.Debug("AddUser")

	// create the empty user model
	user := new(models.User)

	// fill the model with the contents of the request
	err := ctx.Bind(user)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create user from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the user to the database
	user, dbError := server.DB.AddUser(ctx.Request().Context(), user)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding user to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created user
	return ctx.JSON(http.StatusCreated, user)

}
