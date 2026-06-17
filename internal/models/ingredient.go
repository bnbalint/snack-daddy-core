package models

import "time"

type Ingredient struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
