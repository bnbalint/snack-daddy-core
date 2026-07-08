package models

import "time"

type SnackLog struct {
	ID        int
	SnackID   int   // foreign key to the snacks table
	Snack     Snack // gorm will use the ID to populate the object
	TeamID    int   // foreign key to the teams table
	Team      Team  // gorm will use the ID to populate the object
	DateMade  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Override the the default table name used by gorm
func (SnackLog) TableName() string {
	return "snack_log"
}
