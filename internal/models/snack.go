package models

import (
	"time"
)

type Snack struct {
	ID          int
	Name        string
	Sweet       bool
	Savory      bool
	Difficulty  int
	RecipeUrl   string
	Ingredients []Ingredient `gorm:"many2many:snack_ingredients;joinForeignKey:SnackId;joinReferences:IngredientId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
