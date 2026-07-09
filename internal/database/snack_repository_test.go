package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
)

func TestSnackRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  TESTS
	//

	// --- Subtest: Add Snack (Reminder - single DBClient is used across all tests, cannot duplicate test snack) ---
	testingFramework.Run("Add Snack", func(t *testing.T) {
		snack := models.Snack{
			Name:       "Bacon Crackers",
			Sweet:      true,
			Savory:     false,
			Difficulty: 2,
			RecipeUrl:  "",
		}

		savedSnack, err := DbClient.AddSnack(ctx, &snack)
		if err != nil {
			t.Errorf("unexpected error creating snack: %v", err)
		}

		if savedSnack.ID == 0 {
			t.Error("expected snack ID to be populated, got 0")
		}

		fmt.Print("Added snack = ", savedSnack)
	})

	// --- Subtest: Get All Snacks ---
	testingFramework.Run("Get All Snacks", func(t *testing.T) {
		snacks, err := DbClient.GetAllSnacks(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching snacks: %v", err)
		}

		if len(snacks) == 0 {
			t.Errorf("expected some snacks, got '%d'", len(snacks))
		}

		fmt.Print("Retrieved snacks = ", snacks)
	})
}
