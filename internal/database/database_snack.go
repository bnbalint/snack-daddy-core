package database

import (
	"context"
	"log"
	"snack-daddy-core/internal/models"
)

// file for interacting with the snacks table

// get all teams
func (client Client) GetAllSnacks(ctx context.Context) ([]models.Snack, error) {
	var snacks []models.Snack
	result := client.DB.WithContext(ctx).
		Find(&snacks)
	log.Printf("All snacks: %v", snacks)
	return snacks, result.Error
}
