package server

import (
	"fmt"
	"net/http"

	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for rinks

// Get all rinks
// Returns:
//
//	200 and a list of all rinks
func (server *SnackDaddyEchoServer) GetAllRinks(ctx echo.Context) error {
	server.Logger.Debug("GetAllRinks")

	rinks := models.AllRinks()

	server.Logger.Info(fmt.Sprintf("All rinks = %v", rinks))
	return ctx.JSON(http.StatusOK, rinks)
}
