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
	testingFramework.Run("AddUser_sucess", func(t *testing.T) {

		//--------------------------------------------------
		// SET VALUES
		user := models.User{
			FirstName: "Roger",
			LastName:  "Hogwarts",
			Email:     "r.h@gmail.com",
			Teams:     []models.Team{TEAM},
			Allergies: []models.Ingredient{},
		}

		//--------------------------------------------------
		// EXECUTE
		savedUser, err := DbClient.AddUser(ctx, &user)

		//--------------------------------------------------
		// VERIFY RESULTS
		if err != nil {
			t.Errorf("unexpected error creating user: %v", err)
		}

		if savedUser.ID == 0 {
			t.Error("expected user ID to be populated, got 0")
		}
		TEAM.ID = savedUser.Teams[0].ID // save off the created ID

		fmt.Print("Added user = ", savedUser)
	})

	// --- Subtest: GetAllUsers ---
	testingFramework.Run("GetAllUsers_success", func(t *testing.T) {

		//--------------------------------------------------
		// EXECUTE
		users, err := DbClient.GetAllUsers(ctx)

		//--------------------------------------------------
		// VERIFY RESULTS
		if err != nil {
			t.Fatalf("unexpected error fetching users: %v", err)
		}

		if len(users) == 0 {
			t.Errorf("expected some users, got '%d'", len(users))
		}

		fmt.Print("Retrieved users = ", users)
	})

	// --- Subtest: GetUserById_success ---
	testingFramework.Run("GetUserById_success", func(t *testing.T) {

		//--------------------------------------------------
		// SET VALUES
		USER := models.User{
			FirstName: "Roger",
			LastName:  "Hogwarts",
			Email:     "GetUserByIdTest@gmail.com",
			Teams:     []models.Team{},
			Allergies: []models.Ingredient{},
		}

		// first save a user to get the ID
		savedUser, err := DbClient.AddUser(ctx, &USER)

		//--------------------------------------------------
		// EXECUTE
		user, err := DbClient.GetUserById(ctx, savedUser.ID)

		//--------------------------------------------------
		// VERIFY RESULTS

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

	// --- Subtest: GetUserById_notFound ---
	testingFramework.Run("GetUserById_notFound", func(t *testing.T) {

		//--------------------------------------------------
		// EXECUTE
		user, err := DbClient.GetUserById(ctx, 500)

		//--------------------------------------------------
		// VERIFY RESULTS

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

	// --- Subtest: GetAllUsersOnTeam ---
	testingFramework.Run("GetAllUsersOnTeam_success", func(t *testing.T) {

		//--------------------------------------------------
		// EXECUTE
		users, err := DbClient.GetAllUsersOnTeam(ctx, TEAM.ID)

		//--------------------------------------------------
		// VERIFY RESULTS
		if err != nil {
			t.Fatalf("unexpected error fetching users: %v", err)
		}

		if len(users) == 0 {
			t.Errorf("expected some users, got '%d'", len(users))
		}

		fmt.Print("Retrieved users = ", users)
	})

	// --- Subtest: UpdateUser_success ---
	testingFramework.Run("UpdateUser_success", func(t *testing.T) {

		//--------------------------------------------------
		// SET VALUES
		USER := models.User{
			FirstName: "Roger",
			LastName:  "Hogwarts",
			Email:     "UpdateUserTest@gmail.com",
			Teams:     []models.Team{},
			Allergies: []models.Ingredient{},
		}

		// first save a user to get the ID
		savedUser, err := DbClient.AddUser(ctx, &USER)

		// make an update
		savedUser.FirstName = "Updated"

		//--------------------------------------------------
		// EXECUTE
		updatedUser, err := DbClient.UpdateUser(ctx, savedUser)

		//--------------------------------------------------
		// VERIFY RESULTS

		// verify no error
		if err != nil {
			t.Fatalf("unexpected error updating user: %v", err)
		}

		fmt.Print("Updated user = ", updatedUser)
		if updatedUser.FirstName != "Updated" {
			t.Errorf("User name was not updated as expected")
		}
	})
}
