package database

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"
)

// file for interacting with the users table

// get all users
func (client DatabaseClient) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := client.DB.WithContext(ctx).
		Find(&users)
	log.Printf("All users: %v", users)
	return users, result.Error
}

// Add a user to the users table
func (client DatabaseClient) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	result := client.DB.WithContext(ctx).
		Create(&user)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("User created: %v", user)
	return user, nil
}
