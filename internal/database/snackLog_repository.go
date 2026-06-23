package database

import (
	"context"
	"errors"
	"log"
	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"gorm.io/gorm"
)

// file for interacting with the snack_log table

// get all snack_log entries
func (client DatabaseClient) GetSnackLog(ctx context.Context) ([]models.SnackLog, error) {
	var snackLogEntries []models.SnackLog
	result := client.DB.WithContext(ctx).
		Find(&snackLogEntries)
	log.Printf("All snack log entries: %v", snackLogEntries)
	return snackLogEntries, result.Error
}

// Add an entry to the snack log table
func (client DatabaseClient) AddToSnackLog(ctx context.Context, snackLogEntry *models.SnackLog) (*models.SnackLog, error) {
	result := client.DB.WithContext(ctx).
		Create(&snackLogEntry)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("Snack Log entry created: %v", snackLogEntry)
	return snackLogEntry, nil
}
