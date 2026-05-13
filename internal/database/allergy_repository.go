package database

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// file for interacting with the allergies table

// get all allergies
func (client DatabaseClient) GetAllAllergies(ctx context.Context) ([]models.Allergy, error) {
	var allergies []models.Allergy
	result := client.DB.WithContext(ctx).
		Find(&allergies)
	log.Printf("All allergies: %v", allergies)
	return allergies, result.Error
}

// Add an allergy to the allergies table
func (client DatabaseClient) AddAllergy(ctx context.Context, allergy *models.Allergy) (*models.Allergy, error) {
	result := client.DB.WithContext(ctx).
		Create(&allergy)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("Allergy created: %v", allergy)
	return allergy, nil
}
