package models

import "time"

type UserSnackRanking struct {
	SnackID   int   // foreign key to the snacks table
	Snack     Snack // gorm will use the ID to populate the object
	UserID    int   // foreign key to the users table
	User      User  // gorm will use the ID to populate the object
	Rank      SnackRank
	CreatedAt time.Time
	UpdatedAt time.Time
}
