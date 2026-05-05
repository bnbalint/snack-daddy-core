package models

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Teams     []Team    `gorm:"many2many:team_membership;joinForeignKey:UserId;joinReferences:TeamId"`
	Allergies []Allergy `gorm:"many2many:user_allergies;joinForeignKey:UserId;joinReferences:AllergyId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
