package server

import (
	"fmt"
	"net/http"

	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is the controller for levels

// Get all levels
// Returns:
//
//	200 and a list of all levels
func (server *SnackDaddyEchoServer) GetAllLevels(ctx echo.Context) error {
	server.Logger.Debug("GetAllLevels")

	levels := models.AllLevels()

	server.Logger.Info(fmt.Sprintf("All levels = %v", levels))
	return ctx.JSON(http.StatusOK, levels)
}
