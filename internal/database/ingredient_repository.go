package database

import (
	"context"
	"errors"
	"log"
	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"gorm.io/gorm"
)

// file for interacting with the ingredients table

// get all ingredients
func (client DatabaseClient) GetAllIngredients(ctx context.Context) ([]models.Ingredient, error) {
	var ingredients []models.Ingredient
	result := client.DB.WithContext(ctx).
		Find(&ingredients)
	log.Printf("All ingredients: %v", ingredients)
	return ingredients, result.Error
}

// Add an ingredient to the ingredients table
func (client DatabaseClient) AddIngredient(ctx context.Context, ingredient *models.Ingredient) (*models.Ingredient, error) {
	result := client.DB.WithContext(ctx).
		Create(&ingredient)

	if result.Error != nil {

		// if there is a conflict, return our custom error
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &database_errors.ConflictError{}
		}

		// otherwise, return the error as-is
		return nil, result.Error
	}

	log.Printf("Ingredient created: %v", ingredient)
	return ingredient, nil
}
