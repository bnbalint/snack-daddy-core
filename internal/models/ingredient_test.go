package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Run just this test with   go test ingredient_test.go ingredient.go -v
func TestIngredientSerialization(testFramework *testing.T) {

	TIME, _ := time.Parse(time.RFC3339, "2026-07-01T00:00:00Z")

	//--------------------------------------------------
	// SET VALUES
	original := Ingredient{
		ID:        1,
		Name:      "Pecan",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	fmt.Println("Ingredient = ", original)

	//--------------------------------------------------
	// SERIALIZE
	ingredientJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failed to convert ingredient to json: %v", err)
	}
	fmt.Println("Ingredient json = ", string(ingredientJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded Ingredient
	err = json.Unmarshal(ingredientJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert ingredient json back to to object: %v", err)
	}
	fmt.Println("Decoded ingredient = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as the original %+v", decoded, original)
	}
}
