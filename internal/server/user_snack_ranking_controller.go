package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// this is the controller for userSnackRankings

// Get all userSnackRankings
// Returns:
//
//	200 and a list of all userSnackRankings
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllUserSnackRankings(ctx echo.Context) error {
	server.Logger.Debug("GetAllUserSnackRankings")

	userSnackRankings, err := server.DB.GetAllUserSnackRankings(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All userSnackRankings = %v", userSnackRankings))
	return ctx.JSON(http.StatusOK, userSnackRankings)
}

// Add a new userSnackRanking
// Expects the userSnackRanking to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created userSnackRanking object
//	415 if the body cannot be correctly parsed into a userSnackRanking object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddUserSnackRanking(ctx echo.Context) error {
	server.Logger.Debug("AddUserSnackRanking")

	// create the empty userSnackRanking model
	userSnackRanking := new(models.UserSnackRanking)

	// fill the model with the contents of the request
	err := ctx.Bind(userSnackRanking)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create userSnackRanking from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the user to the database
	userSnackRanking, dbError := server.DB.AddUserSnackRanking(ctx.Request().Context(), userSnackRanking)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding userSnackRanking to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created userSnackRanking
	return ctx.JSON(http.StatusCreated, userSnackRanking)

}
