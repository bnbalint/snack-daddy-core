package server

import (
	"fmt"
	"net/http"

	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for ingredients

// Get all ingredients
// Returns:
//
//	200 and a list of all ingredients
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllIngredients(ctx echo.Context) error {
	server.Logger.Debug("GetAllIngredients")

	ingredients, err := server.DB.GetAllIngredients(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All ingredients = %v", ingredients))
	return ctx.JSON(http.StatusOK, ingredients)
}

// Add a new ingredient
// Expects the ingredient to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created ingredint object
//	415 if the body cannot be correctly parsed into a ingredient object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddIngredient(ctx echo.Context) error {
	server.Logger.Debug("AddIngredient")

	// create the empty ingredient model
	ingredient := new(models.Ingredient)

	// fill the model with the contents of the request
	err := ctx.Bind(ingredient)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create ingredient from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the ingredient to the database
	ingredient, dbError := server.DB.AddIngredient(ctx.Request().Context(), ingredient)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding ingredient to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created ingredient
	return ctx.JSON(http.StatusCreated, ingredient)

}
