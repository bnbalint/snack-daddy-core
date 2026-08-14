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

// get a single team by teamId
// if not found, return NotFoundError
func (client DatabaseClient) GetTeamById(ctx context.Context, teamId int) (*models.Team, error) {

	// get the first result by ID
	var team models.Team
	result := client.DB.WithContext(ctx).
		First(&team, teamId)

	if result.Error != nil {

		// no result
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("No team found for teamId %v", teamId)
			return nil, &database_errors.NotFoundError{Entity: "Team", ID: teamId}
		}

		// otherwise, return error as-is
		return nil, result.Error
	}

	log.Printf("Team for teamId %v = %v", teamId, team)
	return &team, result.Error
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

// Update a team, excluding the ID field
func (client DatabaseClient) UpdateTeam(ctx context.Context, team *models.Team) (*models.Team, error) {

	result := client.DB.WithContext(ctx).
		Model(&models.Team{}).
		Omit("ID"). // do not update the ID field
		Where("id = ?", team.ID).
		Updates(team)

	if result.Error != nil {
		log.Printf("Failed to update team: %v", result.Error)

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	// Check if the record actually existed to be updated
	if result.RowsAffected == 0 {
		log.Printf("no team record was updated")
		return nil, nil
	} else {
		log.Printf("Team updated: %v", team)
		return team, nil
	}
}
