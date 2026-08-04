package models

import "time"

type SuggestedAllergy struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
