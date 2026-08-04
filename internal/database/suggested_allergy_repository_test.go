package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
)

func TestSuggestedAllergyRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  TESTS
	//

	// --- Subtest: Add Suggested Allergy ---
	testingFramework.Run("Add Suggested Allergy", func(t *testing.T) {
		suggestion := models.SuggestedAllergy{
			Name: "Pine nut",
		}

		savedSuggestion, err := DbClient.AddSuggestedAllergy(ctx, &suggestion)
		if err != nil {
			t.Errorf("unexpected error creating suggestedAllergy: %v", err)
		}

		if savedSuggestion.ID == 0 {
			t.Error("expected suggestedAllergy ID to be populated, got 0")
		}

		fmt.Print("Added suggestedAllergy = ", savedSuggestion)
	})

	// --- Subtest: Get All Suggested Allergies ---
	testingFramework.Run("Get All Suggested Allergies", func(t *testing.T) {
		suggestedAllergies, err := DbClient.GetAllSuggestedAllergies(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching suggestedAllergies: %v", err)
		}

		if len(suggestedAllergies) == 0 {
			t.Errorf("expected some suggestedAllergies, got '%d'", len(suggestedAllergies))
		}

		fmt.Print("Retrieved suggestedAllergies = ", suggestedAllergies)
	})
}
