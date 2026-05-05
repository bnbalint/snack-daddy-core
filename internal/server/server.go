package server

import (
	. "fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"os"
	"snack-daddy-core/internal/database"
	"snack-daddy-core/internal/models"
)

// this is our microservice
// every function will be definded here
type SnackDaddyCoreService interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	// Teams
	GetAllTeams(ctx echo.Context) error

	// Users
	GetAllUsers(ctx echo.Context) error

	// Snacks
	GetAllSnacks(ctx echo.Context) error

	// Allergies
	GetAllAllergies(ctx echo.Context) error

	// Snack Log
}

// this is our web server
type EchoServer struct {
	echo   *echo.Echo
	DB     database.DatabaseClient
	Logger *slog.Logger
}

func NewEchoServer(db database.DatabaseClient, logger slog.Logger) SnackDaddyCoreService {
	server := &EchoServer{
		echo:   echo.New(),
		DB:     db,
		Logger: &logger,
	}
	server.registerRoutes()
	return server
}

func (server *EchoServer) registerRoutes() {
	server.echo.GET("/readiness", server.Readiness)
	server.echo.GET("/liveness", server.Liveness)

	// teams
	teams := server.echo.Group("/teams")
	teams.GET("", server.GetAllTeams)
	// teams.POST("", server.AddCustomer)
	// teams.GET("/:id", server.GetCustomerById)
	// teams.PUT("/:id", server.UpdateCustomer)
	// teams.DELETE("/:id", server.DeleteCustomer)

	// users
	users := server.echo.Group("/users")
	users.GET("", server.GetAllUsers)

	// snacks
	snacks := server.echo.Group("/snacks")
	snacks.GET("", server.GetAllSnacks)

	// allergies
	allergies := server.echo.Group("/allergies")
	allergies.GET("", server.GetAllAllergies)

}

// Start the server on port [environment var = APP_PORT] or
// port 4242 if the env var is not defined
func (server *EchoServer) Start() error {
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
func (server *EchoServer) Readiness(ctx echo.Context) error {
	ready := server.DB.Ready()
	server.Logger.Info("Readiness reading", "ready", ready)
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

// Determine the Liveness of the server
func (server *EchoServer) Liveness(ctx echo.Context) error {
	server.Logger.Info("Liveness: OK")
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
