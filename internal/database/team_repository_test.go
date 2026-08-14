package database

import (
	"errors"
	"fmt"
	database_errors "snack-daddy-core/internal/database/errors"
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

	// --- Subtest: AddTeam (Reminder - single DBClient is used across all tests, cannot duplicate test team) ---
	testingFramework.Run("AddTeam", func(t *testing.T) {

		//--------------------------------------------------
		// SET VALUES
		team := models.Team{
			Name:           "AddTeamTest",
			Rink:           models.RinkBairel,
			Level:          models.LevelD4,
			PrimaryColor:   "#e03894",
			SecondaryColor: "#3c07b8",
			TernaryColor:   "#08c868",
			LogoUrl:        "",
		}

		//--------------------------------------------------
		// EXECUTE
		savedTeam, err := DbClient.AddTeam(ctx, &team)

		//--------------------------------------------------
		// VERIFY RESULTS
		if err != nil {
			t.Errorf("unexpected error creating team: %v", err)
		}

		if savedTeam.ID == 0 {
			t.Error("expected team ID to be populated, got 0")
		}

		fmt.Print("Added team = ", savedTeam)
	})

	// --- Subtest: GetAllTeams ---
	testingFramework.Run("GetAllTeams", func(t *testing.T) {

		//--------------------------------------------------
		// EXECUTE
		teams, err := DbClient.GetAllTeams(ctx)

		//--------------------------------------------------
		// VERIFY RESULTS
		if err != nil {
			t.Fatalf("unexpected error fetching teams: %v", err)
		}

		if len(teams) == 0 {
			t.Errorf("expected some teams, got '%d'", len(teams))
		}

		fmt.Print("Retrieved teams = ", teams)
	})

	// --- Subtest: GetTeamById_success ---
	testingFramework.Run("GetTeamById_success", func(t *testing.T) {

		//--------------------------------------------------
		// SET VALUES
		team := models.Team{
			Name:           "GetTeamById_Test",
			Rink:           models.RinkBairel,
			Level:          models.LevelD4,
			PrimaryColor:   "#e03894",
			SecondaryColor: "#3c07b8",
			TernaryColor:   "#08c868",
			LogoUrl:        "",
		}

		// first save a team to get the ID
		savedTeam, err := DbClient.AddTeam(ctx, &team)

		//--------------------------------------------------
		// EXECUTE
		foundTeam, err := DbClient.GetTeamById(ctx, savedTeam.ID)

		//--------------------------------------------------
		// VERIFY RESULTS

		// verify no error
		if err != nil {
			t.Fatalf("unexpected error fetching team: %v", err)
		}

		fmt.Print("Retrieved team = ", foundTeam)

		// verify content
		if foundTeam.Rink != "BAIREL" {
			t.Errorf("Team does not have correct rink")
		}
	})

	// --- Subtest: GetTeamById_notFound ---
	testingFramework.Run("GetTeamById_notFound", func(t *testing.T) {

		//--------------------------------------------------
		// EXECUTE
		team, err := DbClient.GetTeamById(ctx, 500)

		//--------------------------------------------------
		// VERIFY RESULTS

		// verify error
		expectedError := &database_errors.NotFoundError{}
		if !errors.As(err, &expectedError) {
			t.Fatalf("Expected NotFoundError but err is: %v", err)
		}

		// verify nil user
		if team != nil {
			t.Fatalf("Expected user to be nil, but user is %v", team)
		}
	})

	// --- Subtest: UpdateTeam (Reminder - single DBClient is used across all tests, cannot duplicate test team) ---
	testingFramework.Run("UpdateTeam", func(t *testing.T) {

		//--------------------------------------------------
		// SET VALUES
		team := models.Team{
			Name:           "UpdateTeamTest",
			Rink:           models.RinkBairel,
			Level:          models.LevelD4,
			PrimaryColor:   "#e03894",
			SecondaryColor: "#3c07b8",
			TernaryColor:   "#08c868",
			LogoUrl:        "",
		}

		//--------------------------------------------------
		// CONFIGURE
		// first add the team
		savedTeam, err := DbClient.AddTeam(ctx, &team)
		if err != nil {
			t.Errorf("unexpected error creating team: %v", err)
		}

		if savedTeam.ID == 0 {
			t.Error("expected team ID to be populated, got 0")
		}

		// update a value
		savedTeam.PrimaryColor = "#000000"

		//--------------------------------------------------
		// EXECUTE
		updatedTeam, err := DbClient.UpdateTeam(ctx, savedTeam)

		//--------------------------------------------------
		// VERIFY RESULTS
		if err != nil {
			t.Errorf("unexpected error updating team: %v", err)
		}

		if updatedTeam.PrimaryColor != "#000000" {
			t.Error("expected Primary Color to be updated")
		}

		fmt.Print("Updated team = ", updatedTeam)
	})
}
