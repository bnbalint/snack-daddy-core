package database

import (
	"context"
	"errors"
	"log"
	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"gorm.io/gorm"
)

// file for interacting with the suggested_allergies table

// get all suggestions
func (client DatabaseClient) GetAllSuggestedAllergies(ctx context.Context) ([]models.SuggestedAllergy, error) {
	var suggestions []models.SuggestedAllergy
	result := client.DB.WithContext(ctx).
		Find(&suggestions)
	log.Printf("All suggestions: %v", suggestions)
	return suggestions, result.Error
}

// Add a suggestion to the suggested_allergies table
func (client DatabaseClient) AddSuggestedAllergy(ctx context.Context, suggestion *models.SuggestedAllergy) (*models.SuggestedAllergy, error) {
	result := client.DB.WithContext(ctx).
		Create(&suggestion)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("SuggestedAllergy created: %v", suggestion)
	return suggestion, nil
}
