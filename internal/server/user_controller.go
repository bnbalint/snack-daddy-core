package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// this is the controller for user

// Get all users
// Returns:
//
//	200 and a list of all teams
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
