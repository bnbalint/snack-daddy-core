package models

import (
	"slices"
	"testing"
)

func TestLevelStringValues(testFramework *testing.T) {

	// define the tests we want to run
	tests := []struct {
		testName      string // name of the test
		level         Level  // the level we want to test
		expectedValue string // gthe string value we expect
	}{
		{"LevelD5", LevelD5, "D5"},
		{"LevelD4", LevelD4, "D4"},
		{"LevelD3", LevelD3, "D3"},
	}

	for _, testData := range tests {
		testFramework.Run(testData.testName, func(testFramework *testing.T) {
			receivedString := testData.level.String()

			if receivedString != testData.expectedValue {
				testFramework.Errorf("Expected %v, got %v", testData.expectedValue, receivedString)
			}
		})
	}
}

func TestAllLevels(testFramework *testing.T) {

	levels := AllLevels()

	// correct number
	if len(levels) != 3 {
		testFramework.Errorf("Expected 3 levels, got %d", len(levels))
	}

	// correct contents
	if !slices.Contains(levels, LevelD5) {
		testFramework.Error("AllLevels does not contain LevelD5")
	}

	if !slices.Contains(levels, LevelD4) {
		testFramework.Error("AllLevels does not contain LevelD4")
	}

	if !slices.Contains(levels, LevelD3) {
		testFramework.Error("AllLevels does not contain LevelD3")
	}

}
