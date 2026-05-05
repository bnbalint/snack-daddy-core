package database

import (
	"context"
	"log"
	"snack-daddy-core/internal/models"
)

// file for interacting with the allergies table

// get all allergies
func (client Client) GetAllAllergies(ctx context.Context) ([]models.Allergy, error) {
	var allergies []models.Allergy
	result := client.DB.WithContext(ctx).
		Find(&allergies)
	log.Printf("All allergies: %v", allergies)
	return allergies, result.Error
}
