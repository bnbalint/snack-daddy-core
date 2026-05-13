package server

import (
	"fmt"
	"net/http"
	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for teams

// Get all teams
// Returns:
//
//	200 and a list of all teams
//	500 for all errors
func (server *SnackDaddyEchoServer) GetAllTeams(ctx echo.Context) error {
	server.Logger.Debug("GetAllTeams")

	teams, err := server.DB.GetAllTeams(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	server.Logger.Info(fmt.Sprintf("All teams = %v", teams))
	return ctx.JSON(http.StatusOK, teams)
}

// Add a new team
// Expects the team to be passed in the body of the request
// Returns:
//
//	201 for successful addition, returning newly created team object
//	415 if the body cannot be correctly parsed into a team object
//	409 for a database key conflict
//	500 for all other errors
func (server *SnackDaddyEchoServer) AddTeam(ctx echo.Context) error {
	server.Logger.Debug("AddTeam")

	// create the empty team model
	team := new(models.Team)

	// fill the model with the contents of the request
	err := ctx.Bind(team)

	// return a 415 if we could not parse the request body
	if err != nil {
		server.Logger.Error("Failed to create team from the provided body")
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	// save the team to the database
	team, dbError := server.DB.AddTeam(ctx.Request().Context(), team)

	// check for error
	if dbError != nil {
		server.Logger.Error("Error encountered while adding team to database")

		// set the status code based on the error
		switch dbError.(type) {
		case *database_errors.ConflictError:
			return ctx.JSON(http.StatusConflict, dbError)

		default:
			return ctx.JSON(http.StatusInternalServerError, dbError)
		}
	}

	// return 201, and the created team
	return ctx.JSON(http.StatusCreated, team)

}
