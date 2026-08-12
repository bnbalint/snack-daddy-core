package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
	"time"
)

func TestUserSnackRankingRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  CONSTANTS - for creating the userSnackRanking
	//
	TIME, _ := time.Parse(time.RFC3339, "2026-07-07T00:00:00Z")

	SNACK := models.Snack{
		Name:        "Ranking Test Treat",
		Sweet:       true,
		Savory:      false,
		Difficulty:  2,
		RecipeUrl:   "",
		Ingredients: []models.Ingredient{},
		CreatedAt:   TIME,
		UpdatedAt:   TIME,
	}

	USER := models.User{
		FirstName: "Roger",
		LastName:  "Hogwarts",
		Email:     "RankingTest@gmail.com",
		Teams:     []models.Team{},
		Allergies: []models.Ingredient{},
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	//---------------------------------------------
	//  TESTS
	//

	// --- Subtest: AddUserSnackRanking ---
	testingFramework.Run("AddUserSnackRanking", func(t *testing.T) {

		userSnackRanking := models.UserSnackRanking{
			Snack:     SNACK,
			User:      USER,
			Rank:      models.SnackRank1,
			CreatedAt: TIME,
			UpdatedAt: TIME,
		}

		savedUserSnackRanking, err := DbClient.AddUserSnackRanking(ctx, &userSnackRanking)
		if err != nil {
			t.Errorf("unexpected error creating userSnackRanking: %v", err)
		}

		if savedUserSnackRanking.SnackID == 0 {
			t.Error("expected saved Snack ID to be populated")
		}
		SNACK.ID = savedUserSnackRanking.SnackID // save this for other tests

		if savedUserSnackRanking.UserID == 0 {
			t.Error("expected saved UserID to be populated")
		}
		USER.ID = savedUserSnackRanking.UserID // save this for other tests

		fmt.Print("Added userSnackRanking = ", savedUserSnackRanking)
	})

	// --- Subtest: GetAllUserSnackRankings ---
	testingFramework.Run("GetAllUserSnackRankings", func(t *testing.T) {
		rankings, err := DbClient.GetAllUserSnackRankings(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching userSnackRankings: %v", err)
		}

		if len(rankings) == 0 {
			t.Errorf("expected some userSnackRankings, got '%d'", len(rankings))
		}

		fmt.Print("Retrieved userSnackRankings = ", rankings)
	})

	// --- Subtest: GetUserSnackRankingsByUserId ---
	testingFramework.Run("GetUserSnackRankingsByUserId", func(t *testing.T) {
		rankings, err := DbClient.GetUserSnackRankingsByUserId(ctx, USER.ID)
		if err != nil {
			t.Fatalf("unexpected error fetching userSnackRankings: %v", err)
		}

		if len(rankings) == 0 {
			t.Errorf("expected some userSnackRankings, got '%d'", len(rankings))
		}

		fmt.Print("Retrieved userSnackRankings = ", rankings)
	})

	// --- Subtest: UpdateUserSnackRankings ---
	testingFramework.Run("UpdateUserSnackRankings", func(t *testing.T) {

		userSnackRanking := models.UserSnackRanking{
			Snack:     SNACK,
			SnackID:   SNACK.ID,
			User:      USER,
			UserID:    USER.ID,
			Rank:      models.SnackRank5,
			CreatedAt: TIME,
			UpdatedAt: TIME,
		}

		rankings := []models.UserSnackRanking{userSnackRanking}

		savedUserSnackRankings, err := DbClient.UpdateUserSnackRankings(ctx, rankings)
		if err != nil {
			t.Errorf("unexpected error updating userSnackRankings: %v", err)
		}

		if len(savedUserSnackRankings) == 0 {
			t.Error("Expected updated UserSnackRankings to be returned, but got length 0")
		}

		fmt.Print("Updated userSnackRanking = ", savedUserSnackRankings)
	})
}
