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

// Update a list of Snacks
func (client DatabaseClient) UpdateSnacks(ctx context.Context, snacks []models.Snack) ([]models.Snack, error) {

	err := client.DB.Transaction(func(transaction *gorm.DB) error {
		for index := range snacks {
			result := transaction.WithContext(ctx).
				Model(&models.Snack{}).
				Omit("ID"). // do not update the ID field
				Where("id = ?", snacks[index].ID).
				Updates(&snacks[index])

			if result.Error != nil {
				return result.Error // rollback the transaction
			}

			// Check if the record actually existed to be updated
			if result.RowsAffected == 0 {
				log.Printf("No Snack record was updated for Id %v", snacks[index].ID)
			} else {
				log.Printf("Snack updated: %v", snacks[index])
			}
		}

		return nil // no error
	})

	if err != nil {
		log.Printf("failed to update snacks: %v", err)

		// if there is a conflict, return our custom error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, err
	}

	return snacks, nil
}
