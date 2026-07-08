package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSnackLogSerialization(testFramework *testing.T) {
	TIME, _ := time.Parse(time.RFC3339, "2026-07-07T00:00:00Z")

	//--------------------------------------------------
	// SET VALUES
	ingredient1 := Ingredient{
		ID:        1,
		Name:      "Rice Crispy Cereal",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient2 := Ingredient{
		ID:        2,
		Name:      "Margarine",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient3 := Ingredient{
		ID:        3,
		Name:      "Marshmallow",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient4 := Ingredient{
		ID:        4,
		Name:      "Vanilla",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	snack := Snack{
		ID:          1,
		Name:        "Rice Crispie Treat",
		Sweet:       true,
		Savory:      false,
		Difficulty:  2,
		RecipeUrl:   "",
		Ingredients: []Ingredient{ingredient1, ingredient2, ingredient3, ingredient4},
		CreatedAt:   TIME,
		UpdatedAt:   TIME,
	}

	team := Team{
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

	original := SnackLog{
		ID:        1,
		SnackID:   1,
		Snack:     snack,
		TeamID:    1,
		Team:      team,
		DateMade:  time.Date(2026, time.June, 23, 0, 0, 0, 0, time.UTC),
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	fmt.Println("SnackLog = ", original)

	//--------------------------------------------------
	// SERIALIZE
	snackLogJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failed to convert snackLog to json: %v", err)
	}
	fmt.Println("SnackLog json = ", string(snackLogJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded SnackLog
	err = json.Unmarshal(snackLogJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert snackLog json back to object: %v", err)
	}
	fmt.Println("Decoded snackLog = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as original %+v", decoded, original)
	}

}
