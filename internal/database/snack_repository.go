package database

import (
	"context"
	"errors"
	"log"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"gorm.io/gorm"
)

// file for interacting with the snacks table

// get all snacks
func (client DatabaseClient) GetAllSnacks(ctx context.Context) ([]models.Snack, error) {
	var snacks []models.Snack
	result := client.DB.WithContext(ctx).
		Find(&snacks)
	log.Printf("All snacks: %v", snacks)
	return snacks, result.Error
}

// Add a snack to the snacks table
func (client DatabaseClient) AddSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error) {
	result := client.DB.WithContext(ctx).
		Create(&snack)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("Snack created: %v", snack)
	return snack, nil
}

// Save a list of snacks to the database
func (client DatabaseClient) UpdateSnacks(ctx context.Context, snacks []models.Snack) ([]models.Snack, error) {
	result := client.DB.WithContext(ctx).
		Save(&snacks)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("Snacks updated: %v", snacks)
	return snacks, nil
}
