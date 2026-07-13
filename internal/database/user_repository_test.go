package database

import (
	"fmt"
	"snack-daddy-core/internal/models"
	"testing"
)

func TestUserRepository(testingFramework *testing.T) {

	// verify that the DbClient was successfully set up in shared_test.go
	if DbClient == nil {
		testingFramework.Fatal("DbClient is not configured")
	}
	testingFramework.Logf("Connecting to shared DbClient at: %s", DbClient)

	//---------------------------------------------
	//  TESTS
	//

	TEAM := models.Team{
		Name:           "UserRepoTestTeam",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#2c54c0",
		SecondaryColor: "#000000",
		TernaryColor:   "#e967b7",
		LogoUrl:        "",
	}

	// --- Subtest: Add User ---
	testingFramework.Run("Add User", func(t *testing.T) {

		user := models.User{
			FirstName: "Roger",
			LastName:  "Hogwarts",
			Email:     "r.h@gmail.com",
			Teams:     []models.Team{TEAM},
			Allergies: []models.Ingredient{},
		}

		savedUser, err := DbClient.AddUser(ctx, &user)
		if err != nil {
			t.Errorf("unexpected error creating user: %v", err)
		}

		if savedUser.ID == 0 {
			t.Error("expected user ID to be populated, got 0")
		}

		fmt.Print("Added user = ", savedUser)
	})

	// --- Subtest: Get All Users ---
	testingFramework.Run("Get All Users", func(t *testing.T) {
		users, err := DbClient.GetAllUsers(ctx)
		if err != nil {
			t.Fatalf("unexpected error fetching users: %v", err)
		}

		if len(users) == 0 {
			t.Errorf("expected some users, got '%d'", len(users))
		}

		fmt.Print("Retrieved users = ", users)
	})
}
