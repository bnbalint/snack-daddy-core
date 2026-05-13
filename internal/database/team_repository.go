package database

import (
	"context"
	"errors"
	"log"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"gorm.io/gorm"
)

// file for interacting with the teams table

// get all teams
func (client DatabaseClient) GetAllTeams(ctx context.Context) ([]models.Team, error) {
	var teams []models.Team
	result := client.DB.WithContext(ctx).
		Find(&teams)

	log.Printf("All teams: %v", teams)
	return teams, result.Error
}

// Add a team to the teams table
func (client DatabaseClient) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	result := client.DB.WithContext(ctx).
		Create(&team)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("Team created: %v", team)
	return team, nil
}
