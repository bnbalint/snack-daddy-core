package server

import (
	"context"

	"snack-daddy-core/internal/models"
)

// mockDB implements SnackDaddyDatabaseClient for testing
type mockDB struct {
	getAllTeamsFunc func(ctx context.Context) ([]models.Team, error)
	addTeamFunc     func(ctx context.Context, team *models.Team) (*models.Team, error)

	getAllSnacksFunc func(ctx context.Context) ([]models.Snack, error)
	addSnackFunc     func(ctx context.Context, snack *models.Snack) (*models.Snack, error)

	getAllAllergiesFunc func(ctx context.Context) ([]models.Allergy, error)
	addAllergyFunc      func(ctx context.Context, snack *models.Allergy) (*models.Allergy, error)

	getAllUsersFunc func(ctx context.Context) ([]models.User, error)
	addUserFunc     func(ctx context.Context, snack *models.User) (*models.User, error)
}

func (mock *mockDB) Ready() bool {
	return true
}
