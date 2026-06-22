package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
)

func TestIngredientRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  TESTS
	//

	// --- Subtest: Add Ingredient ---
	testingFramework.Run("Add Ingredient", func(t *testing.T) {
		ingredient := models.Ingredient{
			Name: "Pecan",
		}

		savedIngredient, err := DbClient.AddIngredient(ctx, &ingredient)
		if err != nil {
			t.Errorf("unexpected error creating ingredient: %v", err)
		}

		if savedIngredient.ID == 0 {
			t.Error("expected ingredient ID to be populated, got 0")
		}

		fmt.Print("Added ingredient = ", savedIngredient)
	})

	// --- Subtest: Get All Ingredients ---
	testingFramework.Run("Get All Ingredients", func(t *testing.T) {
		ingredients, err := DbClient.GetAllIngredients(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching ingredients: %v", err)
		}

		if len(ingredients) == 0 {
			t.Errorf("expected some ingredients, got '%d'", len(ingredients))
		}

		fmt.Print("Retrieved ingredients = ", ingredients)
	})
}
