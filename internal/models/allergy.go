package models

import "time"

type Allergy struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
