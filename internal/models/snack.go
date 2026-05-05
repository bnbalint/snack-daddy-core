package models

import (
	"database/sql"
	"time"
)

type Snack struct {
	ID         int
	Name       string
	Sweet      bool
	Savory     bool
	Difficulty sql.NullInt64
	RecipeUrl  sql.NullString
	Allergies  []Allergy `gorm:"many2many:snack_allergies;joinForeignKey:SnackId;joinReferences:AllergyId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
