package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// this is the controller for snacks

// get all snacks
func (server *EchoServer) GetAllSnacks(ctx echo.Context) error {
	server.Logger.Debug("GetAllSnacks")

	snacks, err := server.DB.GetAllSnacks(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, snacks)
}
