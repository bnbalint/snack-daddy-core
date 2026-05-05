package database

import (
	"context"
	"log"
	"snack-daddy-core/internal/models"
)

// file for interacting with the teams table

// get all teams
func (client Client) GetAllTeams(ctx context.Context) ([]models.Team, error) {
	var teams []models.Team
	result := client.DB.WithContext(ctx).
		Find(&teams)
	log.Printf("All teams: %v", teams)
	return teams, result.Error
}
