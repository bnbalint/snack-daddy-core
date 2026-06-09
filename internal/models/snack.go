package models

import (
	"time"
)

type Snack struct {
	ID         int
	Name       string
	Sweet      bool
	Savory     bool
	Difficulty int
	RecipeUrl  string
	Allergies  []Allergy `gorm:"many2many:snack_allergies;joinForeignKey:SnackId;joinReferences:AllergyId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
