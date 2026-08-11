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

// get a single user by userId
// if not found, retuen NotFoundError
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

// Update a user, excluding the ID and Email fields
func (client DatabaseClient) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {

	result := client.DB.WithContext(ctx).
		Model(&models.User{}).
		Omit("ID", "Email"). // do not update the ID or Email field
		Where("id = ?", user.ID).
		Updates(user)

	if result.Error != nil {
		log.Printf("Failed to update user: %v", result.Error)

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	// Check if the record actually existed to be updated
	if result.RowsAffected == 0 {
		log.Printf("no user record was updated")
		return nil, nil
	} else {
		log.Printf("User updated: %v", user)
		return user, nil
	}
}
