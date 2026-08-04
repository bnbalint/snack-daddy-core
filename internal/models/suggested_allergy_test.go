package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Run just this test with   go test suggested_allergy_test.go suggested_allergy.go -v
func TestSuggestedAllergySerialization(testFramework *testing.T) {

	TIME, _ := time.Parse(time.RFC3339, "2026-07-01T00:00:00Z")

	//--------------------------------------------------
	// SET VALUES
	original := SuggestedAllergy{
		ID:        1,
		Name:      "Pecan",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	fmt.Println("Suggested Allergy = ", original)

	//--------------------------------------------------
	// SERIALIZE
	suggestionJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failed to convert suggestion to json: %v", err)
	}
	fmt.Println("Suggestion json = ", string(suggestionJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded SuggestedAllergy
	err = json.Unmarshal(suggestionJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert suggestion json back to to object: %v", err)
	}
	fmt.Println("Decoded suggestion = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as the original %+v", decoded, original)
	}
}
