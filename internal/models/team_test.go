package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTeamSerialization(testFramework *testing.T) {

	TIME, _ := time.Parse(time.RFC3339, "2026-07-07T00:00:00Z")

	//--------------------------------------------------
	// SET VALUES
	original := Team{
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
	fmt.Println("Team = ", original)

	//--------------------------------------------------
	// SERIALIZE
	teamJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failes to convert team to json: %v", err)
	}
	fmt.Println("Team json = ", string(teamJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded Team
	err = json.Unmarshal(teamJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert team json back to object: %v", err)
	}
	fmt.Println("Decoded team = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as the original %+v", decoded, original)
	}

}
