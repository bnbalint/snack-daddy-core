package models

import (
	"database/sql"
	"time"
)

type Team struct {
	ID             int
	Name           string
	Rink           string
	Level          string
	PrimaryColor   sql.NullString
	SecondaryColor sql.NullString
	TernaryColor   sql.NullString
	LogoUrl        sql.NullString
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
