package models

import (
	"time"
)

type Snack struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"uniqueIndex"` // Tells GORM this field is unique
	Sweet       bool
	Savory      bool
	Difficulty  int
	RecipeUrl   string
	Ingredients []Ingredient `gorm:"many2many:snack_ingredients;joinForeignKey:SnackId;joinReferences:IngredientId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
