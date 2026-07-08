package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSnackSerialization(testFramework *testing.T) {

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

	original := Snack{
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
	fmt.Println("Snack = ", original)

	//--------------------------------------------------
	// SERIALIZE
	snackJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failed to convert snack to json: %v", err)
	}
	fmt.Println("Snack json = ", string(snackJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded Snack
	err = json.Unmarshal(snackJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert snack json back to object: %v", err)
	}
	fmt.Println("Decoded snack = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as original %+v", decoded, original)
	}

}
