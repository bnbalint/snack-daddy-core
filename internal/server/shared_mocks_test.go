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
	updateSnackFunc  func(ctx context.Context, snack *models.Snack) (*models.Snack, error)

	getAllIngredientsFunc func(ctx context.Context) ([]models.Ingredient, error)
	addIngredientsFunc    func(ctx context.Context, snack *models.Ingredient) (*models.Ingredient, error)

	getAllUsersFunc func(ctx context.Context) ([]models.User, error)
	addUserFunc     func(ctx context.Context, snack *models.User) (*models.User, error)

	getSnackLogFunc   func(ctx context.Context) ([]models.SnackLog, error)
	addToSnackLogFunc func(ctx context.Context, snack *models.SnackLog) (*models.SnackLog, error)
}

func (mock *mockDB) Ready() bool {
	return true
}
