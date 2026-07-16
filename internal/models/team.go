package models

import (
	"time"
)

type Team struct {
	ID             int    `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"uniqueIndex"` // Tells GORM this field is unique
	Rink           Rink
	Level          Level
	PrimaryColor   string
	SecondaryColor string
	TernaryColor   string
	LogoUrl        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
