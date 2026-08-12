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
	GetUserById(ctx context.Context, userId int) (*models.User, error)
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, uset *models.User) (*models.User, error)

	// Snacks
	GetAllSnacks(ctx context.Context) ([]models.Snack, error)
	AddSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error)
	UpdateSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error)
	UpdateSnacks(ctx context.Context, snacks []models.Snack) ([]models.Snack, error)

	// Ingredients
	GetAllIngredients(ctx context.Context) ([]models.Ingredient, error)
	AddIngredient(ctx context.Context, ingredient *models.Ingredient) (*models.Ingredient, error)

	// SuggestedAllergies
	GetAllSuggestedAllergies(ctx context.Context) ([]models.SuggestedAllergy, error)
	AddSuggestedAllergy(ctx context.Context, suggestion *models.SuggestedAllergy) (*models.SuggestedAllergy, error)

	// Snack Log
	GetSnackLog(ctx context.Context) ([]models.SnackLog, error)
	AddToSnackLog(ctx context.Context, ingredient *models.SnackLog) (*models.SnackLog, error)

	// Users Snack Rankings
	GetAllUserSnackRankings(ctx context.Context) ([]models.UserSnackRanking, error)
	GetUserSnackRankingsByUserId(ctx context.Context, userId int) ([]models.UserSnackRanking, error)
	AddUserSnackRanking(ctx context.Context, user *models.UserSnackRanking) (*models.UserSnackRanking, error)
	UpdateUserSnackRankings(ctx context.Context, rankings []models.UserSnackRanking) ([]models.UserSnackRanking, error)
}

type DatabaseClient struct {
	DB *gorm.DB
}

// Create a new DatabaseClient using the provided credentials
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
