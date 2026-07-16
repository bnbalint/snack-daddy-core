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
			Sweet:      false,
			Savory:     true,
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

	// --- Subtest: Update Snack ---
	testingFramework.Run("Update Snack - Update Difficulty of Existing Snack", func(t *testing.T) {

		// we can't guarantee what was added in previous tests, so we need to add a snack first, then use the EXACT snack and update it
		snack := models.Snack{
			Name:       "snack to update",
			Sweet:      true,
			Savory:     false,
			Difficulty: 1,
			RecipeUrl:  "",
		}

		// add the snack
		savedSnack, err := DbClient.AddSnack(ctx, &snack)
		if err != nil {
			t.Errorf("unexpected error creating snack: %v", err)
		}
		fmt.Print("Saved snack = ", savedSnack)

		if savedSnack.ID == 0 {
			t.Error("expected snack ID to be populated, got 0")
		}

		// now update the difficulty
		savedSnack.Difficulty = 10
		fmt.Print("Saved snack updated to = ", savedSnack)

		updatedSnack, err := DbClient.UpdateSnack(ctx, &snack)
		if err != nil {
			t.Errorf("unexpected error updating snack: %v", err)
		}
		fmt.Print("Updated snack = ", updatedSnack)

		// check that it was changed
		if updatedSnack.Difficulty != 10 {
			t.Errorf("expected snack difficulty to be 10, got %v", updatedSnack.Difficulty)
		}

	})

}
