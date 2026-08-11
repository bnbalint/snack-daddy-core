package database

import (
	"errors"
	"fmt"
	database_errors "snack-daddy-core/internal/database/errors"
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

	// --- Subtest: Get User By Id ---
	testingFramework.Run("Get User By Id - Success", func(t *testing.T) {

		USER := models.User{
			FirstName: "Roger",
			LastName:  "Hogwarts",
			Email:     "GetUserByIdTest@gmail.com",
			Teams:     []models.Team{},
			Allergies: []models.Ingredient{},
		}

		// first save a user to get the ID
		savedUser, err := DbClient.AddUser(ctx, &USER)

		// get the user by ID
		user, err := DbClient.GetUserById(ctx, savedUser.ID)

		// verify no error
		if err != nil {
			t.Fatalf("unexpected error fetching user: %v", err)
		}

		fmt.Print("Retrieved user = ", user)

		// verify content
		if user.FirstName != "Roger" {
			t.Errorf("User does not have correct first name")
		}
	})

	testingFramework.Run("Get User By Id - Not Found", func(t *testing.T) {

		// get the user by ID
		user, err := DbClient.GetUserById(ctx, 500)

		// verify error
		expectedError := &database_errors.NotFoundError{}
		if !errors.As(err, &expectedError) {
			t.Fatalf("Expected NotFoundError but err is: %v", err)
		}

		// verify nil user
		if user != nil {
			t.Fatalf("Expected user to be nil, but user is %v", user)
		}
	})
}
