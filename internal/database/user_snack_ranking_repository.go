package database

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// file for interacting with the user_snack_rankings table

// get all userSnackRankings
func (client DatabaseClient) GetAllUserSnackRankings(ctx context.Context) ([]models.UserSnackRanking, error) {
	var rankings []models.UserSnackRanking
	result := client.DB.WithContext(ctx).
		Find(&rankings)
	log.Printf("All UserSnackRankings: %v", rankings)
	return rankings, result.Error
}

// get all userSnackRankings for a single user
func (client DatabaseClient) GetUserSnackRankingsByUserId(ctx context.Context, userId int) ([]models.UserSnackRanking, error) {
	var rankings []models.UserSnackRanking
	result := client.DB.WithContext(ctx).
		Where(&models.UserSnackRanking{UserID: userId}).
		Find(&rankings)
	log.Printf("All UserSnackRankings for userId %v: %v", userId, rankings)
	return rankings, result.Error
}

// Add a userSnackRanking to the user_snack_rankings table
func (client DatabaseClient) AddUserSnackRanking(ctx context.Context, userSnackRanking *models.UserSnackRanking) (*models.UserSnackRanking, error) {
	result := client.DB.WithContext(ctx).
		Create(&userSnackRanking)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("UserSnackRanking created: %v", userSnackRanking)
	return userSnackRanking, nil
}

// Update a list of UserSnackRankings
func (client DatabaseClient) UpdateUserSnackRankings(ctx context.Context, rankings []models.UserSnackRanking) ([]models.UserSnackRanking, error) {

	err := client.DB.Transaction(func(transaction *gorm.DB) error {
		for index := range rankings {
			result := transaction.WithContext(ctx).
				Model(&models.UserSnackRanking{}).
				Omit("SnackID", "UserID"). // do not update the ID field
				Where("snack_id = ? AND user_id = ?", rankings[index].SnackID, rankings[index].UserID).
				Updates(&rankings[index])

			if result.Error != nil {
				return result.Error // rollback the transaction
			}

			// Check if the record actually existed to be updated
			if result.RowsAffected == 0 {
				log.Printf("No UserSnackRanking record was updated for snackId %v and userId %v", rankings[index].SnackID, rankings[index].UserID)
			} else {
				log.Printf("UserSnackRanking updated: %v", rankings[index])
			}
		}

		return nil // no error
	})

	if err != nil {
		log.Printf("failed to update rankings: %v", err)

		// if there is a conflict, return our custom error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, err
	}

	return rankings, nil
}
