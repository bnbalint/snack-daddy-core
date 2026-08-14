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

	// --- Subtest: AddSnack ---
	testingFramework.Run("AddSnack", func(t *testing.T) {
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

	// --- Subtest: GetAllSnacks ---
	testingFramework.Run("GetAllSnacks", func(t *testing.T) {
		snacks, err := DbClient.GetAllSnacks(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching snacks: %v", err)
		}

		if len(snacks) == 0 {
			t.Errorf("expected some snacks, got '%d'", len(snacks))
		}

		fmt.Print("Retrieved snacks = ", snacks)
	})

	// --- Subtest: UpdateSnacks ---
	testingFramework.Run("UpdateSnacks", func(t *testing.T) {

		// we can't guarantee what was added in previous tests, so we need to add a snack first, then use the EXACT snack and update it
		snack1 := models.Snack{
			Name:       "Multi-Update",
			Sweet:      true,
			Savory:     false,
			Difficulty: 1,
			RecipeUrl:  "",
		}

		snack2 := models.Snack{
			Name:       "Multi-Update2",
			Sweet:      true,
			Savory:     false,
			Difficulty: 1,
			RecipeUrl:  "",
		}

		// add the snacks
		savedSnack, err := DbClient.AddSnack(ctx, &snack1)
		if err != nil {
			t.Errorf("unexpected error creating snack: %v", err)
		}
		fmt.Print("Saved snack = ", savedSnack)

		if savedSnack.ID == 0 {
			t.Error("expected snackID to be populated, got 0")
		}

		savedSnack2, err := DbClient.AddSnack(ctx, &snack2)
		if err != nil {
			t.Errorf("unexpected error creating snack: %v", err)
		}
		fmt.Print("Saved snack = ", savedSnack)

		if savedSnack2.ID == 0 {
			t.Error("expected snackID to be populated, got 0")
		}

		// now update the difficulty (for one - this mimics real life usage where not all snacks will necessarily have changes)
		savedSnack.Difficulty = 10
		fmt.Print("Saved snack updated to = ", savedSnack)

		snacks := []models.Snack{*savedSnack, *savedSnack2}

		updatedSnacks, err := DbClient.UpdateSnacks(ctx, snacks)
		if err != nil {
			t.Errorf("unexpected error updating snacks: %v", err)
		}
		fmt.Print("Updated snacks = ", updatedSnacks)

		// check size
		if len(updatedSnacks) != 2 {
			t.Errorf("Expected 2 snacks to be returned, got %v", len(updatedSnacks))
		}

		// check that it was changed
		if updatedSnacks[0].Difficulty != 10 {
			t.Errorf("expected snack difficulty to be 10, got %v", updatedSnacks[0].Difficulty)
		}

	})

}
