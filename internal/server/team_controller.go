package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// this is the controller for teams

// get all teams
func (server *EchoServer) GetAllTeams(ctx echo.Context) error {
	server.Logger.Debug("GetAllTeams")

	teams, err := server.DB.GetAllTeams(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, teams)
}
