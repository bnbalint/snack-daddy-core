package server

import (
	"context"

	"snack-daddy-core/internal/models"
)

// mockDB implements SnackDaddyDatabaseClient for testing
type mockDB struct {

	// Teams
	getAllTeamsFunc func(ctx context.Context) ([]models.Team, error)
	addTeamFunc     func(ctx context.Context, team *models.Team) (*models.Team, error)
	updateTeamFunc  func(ctx context.Context, team *models.Team) (*models.Team, error)

	// Snacks
	getAllSnacksFunc func(ctx context.Context) ([]models.Snack, error)
	addSnackFunc     func(ctx context.Context, snack *models.Snack) (*models.Snack, error)
	updateSnacksFunc func(ctx context.Context, snacks []models.Snack) ([]models.Snack, error)

	// Ingredients
	getAllIngredientsFunc func(ctx context.Context) ([]models.Ingredient, error)
	addIngredientsFunc    func(ctx context.Context, snack *models.Ingredient) (*models.Ingredient, error)

	// SuggestedAllergies
	getAllSuggestedAllergiesFunc func(ctx context.Context) ([]models.SuggestedAllergy, error)
	addSuggestedAllergyFunc      func(ctx context.Context, snack *models.SuggestedAllergy) (*models.SuggestedAllergy, error)

	// Users
	getAllUsersFunc       func(ctx context.Context) ([]models.User, error)
	getUserByIdFunc       func(ctx context.Context, userId int) (*models.User, error)
	getAllUsersOnTeamFunc func(ctx context.Context, teamId int) ([]models.User, error)
	addUserFunc           func(ctx context.Context, snack *models.User) (*models.User, error)
	updateUserFunc        func(ctx context.Context, user *models.User) (*models.User, error)

	// SnackLog
	getSnackLogFunc   func(ctx context.Context) ([]models.SnackLog, error)
	addToSnackLogFunc func(ctx context.Context, snack *models.SnackLog) (*models.SnackLog, error)

	// UserSnackRankings
	getAllUserSnackRankingsFunc      func(ctx context.Context) ([]models.UserSnackRanking, error)
	getUserSnackRankingsByUserIdFunc func(ctx context.Context, userId int) ([]models.UserSnackRanking, error)
	addUserSnackRankingFunc          func(ctx context.Context, snack *models.UserSnackRanking) (*models.UserSnackRanking, error)
	updateUserSnackRankingsFunc      func(ctx context.Context, rankings []models.UserSnackRanking) ([]models.UserSnackRanking, error)
}

func (mock *mockDB) Ready() bool {
	return true
}
