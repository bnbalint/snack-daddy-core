package models

import "time"

type Ingredient struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"uniqueIndex"` // Tells GORM this field is unique
	CreatedAt time.Time
	UpdatedAt time.Time
}
