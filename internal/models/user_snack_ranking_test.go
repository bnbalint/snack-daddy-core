package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestUserSnackRankingSerialization(testFramework *testing.T) {
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

	SNACK := Snack{
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

	PECAN := Ingredient{
		ID:        1,
		Name:      "Pecan",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	TEAM_MULES := Team{
		ID:             1,
		Name:           "Mules",
		Rink:           RinkBairel,
		Level:          LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
		CreatedAt:      TIME,
		UpdatedAt:      TIME,
	}

	USER := User{
		ID:        1,
		FirstName: "Roger",
		LastName:  "Hogwarts",
		Email:     "r.h@gmail.com",
		Teams:     []Team{TEAM_MULES},
		Allergies: []Ingredient{PECAN},
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	original := UserSnackRanking{
		SnackID:   1,
		Snack:     SNACK,
		UserID:    1,
		User:      USER,
		Rank:      SnackRank1,
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	fmt.Println("UserSnackRanking = ", original)

	//--------------------------------------------------
	// SERIALIZE
	snackRankingJson, err := json.Marshal(original)
	if err != nil {
		testFramework.Fatalf("Failed to convert UserSnackRanking to json: %v", err)
	}
	fmt.Println("UserSnackRanking json = ", string(snackRankingJson))

	//--------------------------------------------------
	// DESERIALIZE
	var decoded UserSnackRanking
	err = json.Unmarshal(snackRankingJson, &decoded)
	if err != nil {
		testFramework.Fatalf("Failed to convert snackRanking json back to object: %v", err)
	}
	fmt.Println("Decoded snackRanking = ", decoded)

	//--------------------------------------------------
	// VERIFY RESULTS
	if !reflect.DeepEqual(original, decoded) {
		testFramework.Errorf("Decoded %+v is not the same as original %+v", decoded, original)
	}

}
