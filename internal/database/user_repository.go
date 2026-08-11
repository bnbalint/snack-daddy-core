package database

import (
	"context"
	"errors"
	"log"
	"snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"gorm.io/gorm"
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

// get a single user
func (client DatabaseClient) GetUserById(ctx context.Context, userId int) (*models.User, error) {

	// get the first result by ID
	var user models.User
	result := client.DB.WithContext(ctx).
		First(&user, userId)

	if result.Error != nil {

		// no result
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("No user found for userId %v", userId)
			return nil, &database_errors.NotFoundError{Entity: "User", ID: userId}
		}

		// otherwise, return error as-is
		return nil, result.Error
	}

	log.Printf("User for userId %v = %v", userId, user)
	return &user, result.Error
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
