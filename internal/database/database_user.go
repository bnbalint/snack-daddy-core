package database

import (
	"context"
	"log"
	"snack-daddy-core/internal/models"
)

// file for interacting with the users table

// get all users
func (client Client) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := client.DB.WithContext(ctx).
		Find(&users)
	log.Printf("All users: %v", users)
	return users, result.Error
}
