package server

import (
	"fmt"
	"net/http"

	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for the snacklog

// Get the snack log
// Returns:
//
//	200 and the entire snack log
//	500 for all errors
func (server *SnackDaddyEchoServer) GetSnackLog(ctx echo.Context) error {
	server.Logger.Debug("GetSnackLog")

	snackLog, err := server.DB.GetSnackLog(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("Snack Log = %v", snackLog))
	return ctx.JSON(http.StatusOK, snackLog)
}

// Add an entry to the snack log
// Expects the entry to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created entry
//	415 if the body cannot be correctly parsed into a snack log entry
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddToSnackLog(ctx echo.Context) error {
	server.Logger.Debug("AddToSnackLog")

	// create the empty SnackLog
	snackLog := new(models.SnackLog)

	// fill the model with the contents of the request
	err := ctx.Bind(snackLog)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create snackLog entry from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the snack log entry to the database
	snackLog, dbError := server.DB.AddToSnackLog(ctx.Request().Context(), snackLog)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding snack log to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created entry
	return ctx.JSON(http.StatusCreated, snackLog)

}
