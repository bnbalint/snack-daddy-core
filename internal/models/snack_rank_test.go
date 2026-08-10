package models

import (
	"slices"
	"testing"
)

func TestSnackRankStringValues(testFramework *testing.T) {

	// define the tests we want to run
	tests := []struct {
		testName      string    // name of the test
		rank          SnackRank // the level we want to test
		expectedValue string    // gthe string value we expect
	}{
		{"SnackRank1", SnackRank1, "RANK_1"},
		{"SnackRank2", SnackRank2, "RANK_2"},
		{"SnackRank3", SnackRank3, "RANK_3"},
		{"SnackRank4", SnackRank4, "RANK_4"},
		{"SnackRank5", SnackRank5, "RANK_5"},
		{"SnackRank6", SnackRank6, "RANK_6"},
		{"SnackRank7", SnackRank7, "RANK_7"},
		{"SnackRank8", SnackRank8, "RANK_8"},
		{"SnackRank9", SnackRank9, "RANK_9"},
		{"SnackRank10", SnackRank10, "RANK_10"},
		{"HaveNotTried", SnackRankHaveNotTried, "HAVE NOT TRIED"},
		{"Unranked", SnackRankUnranked, "UNRANKED"},
	}

	for _, testData := range tests {
		testFramework.Run(testData.testName, func(testFramework *testing.T) {
			receivedString := testData.rank.String()

			if receivedString != testData.expectedValue {
				testFramework.Errorf("Expected %v, got %v", testData.expectedValue, receivedString)
			}
		})
	}
}

func TestAllSnackRanks(testFramework *testing.T) {

	ranks := AllSnackRanks()

	// correct number
	if len(ranks) != 12 {
		testFramework.Errorf("Expected 12 snack ranks, got %d", len(ranks))
	}

	// correct contents
	if !slices.Contains(ranks, SnackRank1) {
		testFramework.Error("AllSnackRanks does not contain SnackRank1")
	}
	if !slices.Contains(ranks, SnackRank2) {
		testFramework.Error("AllSnackRanks does not contain SnackRank2")
	}
	if !slices.Contains(ranks, SnackRank3) {
		testFramework.Error("AllSnackRanks does not contain SnackRank3")
	}
	if !slices.Contains(ranks, SnackRank4) {
		testFramework.Error("AllSnackRanks does not contain SnackRank4")
	}
	if !slices.Contains(ranks, SnackRank5) {
		testFramework.Error("AllSnackRanks does not contain SnackRank5")
	}
	if !slices.Contains(ranks, SnackRank6) {
		testFramework.Error("AllSnackRanks does not contain SnackRank6")
	}
	if !slices.Contains(ranks, SnackRank7) {
		testFramework.Error("AllSnackRanks does not contain SnackRank7")
	}
	if !slices.Contains(ranks, SnackRank8) {
		testFramework.Error("AllSnackRanks does not contain SnackRank8")
	}
	if !slices.Contains(ranks, SnackRank9) {
		testFramework.Error("AllSnackRanks does not contain SnackRank9")
	}
	if !slices.Contains(ranks, SnackRank10) {
		testFramework.Error("AllSnackRanks does not contain SnackRank10")
	}
	if !slices.Contains(ranks, SnackRankHaveNotTried) {
		testFramework.Error("AllSnackRanks does not contain SnackRankHaveNotTried")
	}
	if !slices.Contains(ranks, SnackRankUnranked) {
		testFramework.Error("AllSnackRanks does not contain SnackRankUnranked")
	}

}
