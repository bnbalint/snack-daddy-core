package server

import (
	. "fmt"
	"log/slog"
	"net/http"
	"os"
	"snack-daddy-core/internal/database"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// this is our microservice
// every function will be definded here
type SnackDaddyCoreService interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	// Teams
	GetAllTeams(ctx echo.Context) error
	AddTeam(ctx echo.Context) error

	// Users
	GetAllUsers(ctx echo.Context) error
	AddUser(ctx echo.Context) error

	// Snacks
	GetAllSnacks(ctx echo.Context) error
	AddSnack(ctx echo.Context) error

	// Ingredients
	GetAllIngredients(ctx echo.Context) error
	AddIngredient(ctx echo.Context) error

	// Snack Log
	GetSnackLog(ctx echo.Context) error
	AddToSnackLog(ctx echo.Context) error
}

// this is our web server
type SnackDaddyEchoServer struct {
	echo   *echo.Echo
	DB     database.SnackDaddyDatabaseClient
	Logger *slog.Logger
}

func NewSnackDaddyEchoServer(db database.SnackDaddyDatabaseClient, logger slog.Logger) SnackDaddyCoreService {
	server := &SnackDaddyEchoServer{
		echo:   echo.New(),
		DB:     db,
		Logger: &logger,
	}
	server.registerRoutes()
	return server
}

func (server *SnackDaddyEchoServer) registerRoutes() {
	server.echo.GET("/readiness", server.Readiness)
	server.echo.GET("/liveness", server.Liveness)

	// teams
	teams := server.echo.Group("/teams")
	teams.GET("", server.GetAllTeams)
	teams.POST("", server.AddTeam)
	// teams.GET("/:id", server.GetCustomerById)
	// teams.PUT("/:id", server.UpdateCustomer)
	// teams.DELETE("/:id", server.DeleteCustomer)

	// users
	users := server.echo.Group("/users")
	users.GET("", server.GetAllUsers)
	users.POST("", server.AddUser)

	// snacks
	snacks := server.echo.Group("/snacks")
	snacks.GET("", server.GetAllSnacks)
	snacks.POST("", server.AddSnack)

	// ingredients
	ingredients := server.echo.Group("/ingredients")
	ingredients.GET("", server.GetAllIngredients)
	ingredients.POST("", server.AddIngredient)

	// snacklog
	snackLog := server.echo.Group("/snack-log")
	snackLog.GET("", server.GetSnackLog)
	snackLog.POST("", server.AddToSnackLog)

}

// Start the server on port [environment var = APP_PORT] or
// port 4242 if the env var is not defined
func (server *SnackDaddyEchoServer) Start() error {
	port := os.Getenv("app_port")
	if port == "" {
		server.Logger.Warn("APP_PORT not defined, will use default port")
		port = ":4242"
	}
	server.Logger.Info(Sprintf("Starting server on port %s", port))
	if err := server.echo.Start(port); err != nil && err != http.ErrServerClosed {
		server.Logger.Error("server shutdown occurred: %s", "error", err)
		return err
	}
	return nil
}

// Determines the Readiness of server
// Calls DB.Ready()
//
//	--> 200 if ready
//	--> 500 if not ready
func (server *SnackDaddyEchoServer) Readiness(ctx echo.Context) error {
	ready := server.DB.Ready()
	server.Logger.Info("Readiness reading", "ready", ready)
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

// Determine the Liveness of the server
func (server *SnackDaddyEchoServer) Liveness(ctx echo.Context) error {
	server.Logger.Info("Liveness: OK")
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
