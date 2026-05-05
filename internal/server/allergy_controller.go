package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// this is the controller for allergies

// get all allergies
func (server *EchoServer) GetAllAllergies(ctx echo.Context) error {
	server.Logger.Debug("GetAllAllergies")

	allergies, err := server.DB.GetAllAllergies(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, allergies)
}
