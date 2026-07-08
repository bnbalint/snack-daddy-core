package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestUserSerialization(testFramework *testing.T) {

	TIME, _ := time.Parse(time.RFC3339, "2026-07-07T00:00:00Z")

	//--------------------------------------------------
	// SET VALUES
	PECAN := Ingredient{
		ID:        1,
		Name:      "Pecan",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	TEAM_MULES := Team{
		ID:             1,
		Name:           "Mules",
		Rink:           "BAIREL",
		Level:          "D5",
		PrimaryColor:   "Gold",
		SecondaryColor: "Black",
		TernaryColor:   "Brick Red",
		LogoUrl:        "",
		CreatedAt:      TIME,
		UpdatedAt:      TIME,
	}

	original := User{
		ID:        1,
		FirstName: "Roger",
		LastName:  "Hogwarts",
		Email:     "r.h@gmail.com",
		Teams:     []Team{TEAM_MULES},
		Allergies: []Ingredient{PECAN},
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	fmt.Println("User = ", original)

	//--------------------------------------------------
	// SERIALIZE
	userJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failed to convert user to json: %v", err)
	}
	fmt.Println("User json = ", string(userJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded User
	err = json.Unmarshal(userJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert user json back to object: %v", err)
	}
	fmt.Println("Decoded user = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as the original %+v", decoded, original)
	}
}
