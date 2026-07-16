package models

import "time"

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Email     string       `gorm:"uniqueIndex"` // Tells GORM this field is unique
	Teams     []Team       `gorm:"many2many:team_membership;joinForeignKey:UserId;joinReferences:TeamId"`
	Allergies []Ingredient `gorm:"many2many:user_allergies;joinForeignKey:UserId;joinReferences:IngredientId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
