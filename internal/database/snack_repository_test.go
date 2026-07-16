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

	// --- Subtest: Add Snack ---
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

	// --- Subtest: Update Snacks ---
	testingFramework.Run("Update Snacks", func(t *testing.T) {
		snack1 := models.Snack{
			Name:       "Bacon Crackers",
			Sweet:      false,
			Savory:     true,
			Difficulty: 2,
			RecipeUrl:  "",
		}

		snack2 := models.Snack{
			Name:       "Rice Crispie Treats",
			Sweet:      true,
			Savory:     false,
			Difficulty: 2,
			RecipeUrl:  "",
		}

		snacks := []models.Snack{snack1, snack2}

		updatedSnacks, err := DbClient.UpdateSnacks(ctx, snacks)
		if err != nil {
			t.Errorf("unexpected error updating snacks: %v", err)
		}

		if updatedSnacks[0].ID == 0 {
			t.Error("expected snack ID to be populated, got 0")
		}

		if updatedSnacks[1].ID == 0 {
			t.Error("expected snack ID to be populated, got 0")
		}

		fmt.Print("Updated snacks = ", updatedSnacks)

		// update the difficulty and do the update again
		snack1.Difficulty = 10
		updatedSnacks, err = DbClient.UpdateSnacks(ctx, snacks)
		if err != nil {
			t.Errorf("unexpected error updating snacks: %v", err)
		}

		// check that it was changed
		if updatedSnacks[0].Difficulty != 10 {
			t.Errorf("expected snack difficulty to be 10, got %v", updatedSnacks[0].Difficulty)
		}

	})
}
