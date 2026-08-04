package server

import (
	"fmt"
	"net/http"

	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for suggestedAllergies

// Get all suggestedAllergies
// Returns:
//
//	200 and a list of all suggestedAllergies
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllSuggestedAllergies(ctx echo.Context) error {
	server.Logger.Debug("GetAllSuggestedAllergies")

	suggestedAllergies, err := server.DB.GetAllSuggestedAllergies(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All suggestedAllergies = %v", suggestedAllergies))
	return ctx.JSON(http.StatusOK, suggestedAllergies)
}

// Add a new suggestedAllergy
// Expects the suggestedAllergy to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created suggestedAllergy object
//	415 if the body cannot be correctly parsed into a suggestedAllergy object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddSuggestedAllergy(ctx echo.Context) error {
	server.Logger.Debug("AddSuggestedAllergy")

	// create the empty suggestedAllergy model
	suggestedAllergy := new(models.SuggestedAllergy)

	// fill the model with the contents of the request
	err := ctx.Bind(suggestedAllergy)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create suggestedAllergy from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// TODO - validate the name field

	// save the suggestedAllergy to the database
	suggestedAllergy, dbError := server.DB.AddSuggestedAllergy(ctx.Request().Context(), suggestedAllergy)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding suggestedAllergy to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created ingredient
	return ctx.JSON(http.StatusCreated, suggestedAllergy)

}
