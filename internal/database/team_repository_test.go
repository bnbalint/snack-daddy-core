package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
)

func TestTeamRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  TESTS
	//

	// --- Subtest: Add Team (Reminder - single DBClient is used across all tests, cannot duplicate test team) ---
	testingFramework.Run("Add Team", func(t *testing.T) {
		team := models.Team{
			Name:           "Monsters",
			Rink:           "BAIREL",
			Level:          "D4",
			PrimaryColor:   "#e03894",
			SecondaryColor: "#3c07b8",
			TernaryColor:   "#08c868",
			LogoUrl:        "",
		}

		savedTeam, err := DbClient.AddTeam(ctx, &team)
		if err != nil {
			t.Errorf("unexpected error creating team: %v", err)
		}

		if savedTeam.ID == 0 {
			t.Error("expected team ID to be populated, got 0")
		}

		fmt.Print("Added team = ", savedTeam)
	})

	// --- Subtest: Get All Teams ---
	testingFramework.Run("Get All Teams", func(t *testing.T) {
		teams, err := DbClient.GetAllTeams(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching teams: %v", err)
		}

		if len(teams) == 0 {
			t.Errorf("expected some teams, got '%d'", len(teams))
		}

		fmt.Print("Retrieved teams = ", teams)
	})
}
