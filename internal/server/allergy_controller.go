package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// this is the controller for allergies

// Get all allergies
// Returns:
//
//	200 and a list of all teams
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllAllergies(ctx echo.Context) error {
	server.Logger.Debug("GetAllAllergies")

	allergies, err := server.DB.GetAllAllergies(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All allergies = %v", allergies))
	return ctx.JSON(http.StatusOK, allergies)
}

// Add a new allergy
// Expects the allergy to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created allergy object
//	415 if the body cannot be correctly parsed into a allergy object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddAllergy(ctx echo.Context) error {
	server.Logger.Debug("AddAllergy")

	// create the empty allergy model
	allergy := new(models.Allergy)

	// fill the model with the contents of the request
	err := ctx.Bind(allergy)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create allergy from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the allergy to the database
	allergy, dbError := server.DB.AddAllergy(ctx.Request().Context(), allergy)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding allergy to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created allergy
	return ctx.JSON(http.StatusCreated, allergy)

}
