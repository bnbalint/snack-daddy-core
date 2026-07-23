package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// this is the controller for snacks

// Get all snacks
// Returns:
//
//	200 and a list of all teams
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllSnacks(ctx echo.Context) error {
	server.Logger.Debug("GetAllSnacks")

	snacks, err := server.DB.GetAllSnacks(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All snacks = %v", snacks))
	return ctx.JSON(http.StatusOK, snacks)
}

// Add a new snack
// Expects the snack to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created snack object
//	415 if the body cannot be correctly parsed into a snack object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddSnack(ctx echo.Context) error {
	server.Logger.Debug("AddSnack")

	// create the empty snack model
	snack := new(models.Snack)

	// fill the model with the contents of the request
	err := ctx.Bind(snack)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create snack from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the snack to the database
	snack, dbError := server.DB.AddSnack(ctx.Request().Context(), snack)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding snack to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created snack
	return ctx.JSON(http.StatusCreated, snack)

}

// Update an existing snack
// Expects the snack to be passed in the body of the request
// Returns:
//
//	200 for successful update, returning newly updated snack object
//	415 if the body cannot be correctly parsed into a snack object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) UpdateSnack(ctx echo.Context) error {
	server.Logger.Debug("UpdateSnack")

	// create the empty snack model
	snack := new(models.Snack)

	// fill the model with the contents of the request
	err := ctx.Bind(snack)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create snack from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the snack to the database
	snack, dbError := server.DB.UpdateSnack(ctx.Request().Context(), snack)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while updating snack in the database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 200, and the updated snack
	return ctx.JSON(http.StatusOK, snack)

}
