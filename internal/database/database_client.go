package database

import (
	"context"
	"fmt"
	"snack-daddy-core/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Client for talking to the SnackDaddy database
type SnackDaddyDatabaseClient interface {
	Ready() bool

	// Teams
	GetAllTeams(ctx context.Context) ([]models.Team, error)
	AddTeam(ctx context.Context, team *models.Team) (*models.Team, error)

	// Users
	GetAllUsers(ctx context.Context) ([]models.User, error)
	AddUser(ctx context.Context, user *models.User) (*models.User, error)

	// Snacks
	GetAllSnacks(ctx context.Context) ([]models.Snack, error)
	AddSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error)

	// Ingredients
	GetAllIngredients(ctx context.Context) ([]models.Ingredient, error)
	AddIngredient(ctx context.Context, ingredient *models.Ingredient) (*models.Ingredient, error)

	// Snack Log
}

type DatabaseClient struct {
	DB *gorm.DB
}

// Create a new DatabaseClient using the provided credentials
// TODO: Currently EVIL and will print credentials
func NewDatabaseClient(host string, user string, password string, dbname string, port int32, sslmode string) (SnackDaddyDatabaseClient, error) {

	// collect the connection information into a single string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)

	// create the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		return nil, err
	}

	// create the client
	client := DatabaseClient{DB: db}

	return client, nil
}

// Determine if the databsase is ready
// Performs a basic SELECT statement to determine readiness
func (client DatabaseClient) Ready() bool {
	var ready string
	tx := client.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
