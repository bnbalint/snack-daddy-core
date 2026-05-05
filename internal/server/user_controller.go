package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// this is the controller for user

// get all users
func (server *EchoServer) GetAllUsers(ctx echo.Context) error {
	server.Logger.Debug("GetAllUsers")

	users, err := server.DB.GetAllUsers(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, users)
}
