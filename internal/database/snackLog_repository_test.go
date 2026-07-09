package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
	"time"
)

func TestSnackLogRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  TESTS
	//

	SNACK := models.Snack{
		Name:       "Rice Crispie Treat",
		Sweet:      true,
		Savory:     false,
		Difficulty: 2,
		RecipeUrl:  "",
	}

	TEAM := models.Team{
		Name:           "Mules",
		Rink:           "BAIREL",
		Level:          "D5",
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
	}

	// --- Subtest: Add To SnackLog ---
	testingFramework.Run("Add To Snack Log", func(t *testing.T) {
		snackLogEntry := models.SnackLog{
			SnackID:  1,
			TeamID:   1,
			DateMade: time.Date(2026, time.June, 23, 0, 0, 0, 0, time.UTC),
		}

		// first we need to make sure a team and snack have been added to the test database
		// this is dependent on the order that the test run in
		DbClient.AddSnack(ctx, &SNACK)
		DbClient.AddTeam(ctx, &TEAM)

		savedSnackLogEntry, err := DbClient.AddToSnackLog(ctx, &snackLogEntry)
		if err != nil {
			t.Errorf("unexpected error creating snack log entry: %v", err)
		}

		if savedSnackLogEntry.ID == 0 {
			t.Error("expected snack log entry ID to be populated, got 0")
		}

		fmt.Print("Added snack log entry = ", savedSnackLogEntry)
	})

	// --- Subtest: Get Snack Log ---
	testingFramework.Run("Get Snack Log", func(t *testing.T) {
		snackLog, err := DbClient.GetSnackLog(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching snack log: %v", err)
		}

		if len(snackLog) == 0 {
			t.Errorf("expected some snack log entries, got '%d'", len(snackLog))
		}

		fmt.Print("Retrieved snack log = ", snackLog)
	})
}
