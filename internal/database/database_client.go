package database

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"snack-daddy-core/internal/models"
)

// general client for talking to database

type DatabaseClient interface {
	Ready() bool

	// Teams
	GetAllTeams(ctx context.Context) ([]models.Team, error)

	// Users
	GetAllUsers(ctx context.Context) ([]models.User, error)

	// Snacks
	GetAllSnacks(ctx context.Context) ([]models.Snack, error)

	// Allergies
	GetAllAllergies(ctx context.Context) ([]models.Allergy, error)

	// Snack Log
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient(host string, user string, password string, dbname string, port int32, sslmode string) (DatabaseClient, error) {

	// collect the connection information into a single string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)

	// create the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		return nil, err
	}

	// create the client
	client := Client{DB: db}

	return client, nil
}

func (client Client) Ready() bool {
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
