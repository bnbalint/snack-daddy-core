package models

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Teams     []Team       `gorm:"many2many:team_membership;joinForeignKey:UserId;joinReferences:TeamId"`
	Allergies []Ingredient `gorm:"many2many:user_allergies;joinForeignKey:UserId;joinReferences:IngredientId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
