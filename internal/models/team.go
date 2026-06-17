package models

import (
	"time"
)

type Team struct {
	ID             int
	Name           string
	Rink           string
	Level          string
	PrimaryColor   string
	SecondaryColor string
	TernaryColor   string
	LogoUrl        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
