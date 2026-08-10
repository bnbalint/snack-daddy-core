package models

import (
	"slices"
	"testing"
)

func TestRinkStringValues(testFramework *testing.T) {

	// define the tests we want to run
	tests := []struct {
		testName      string // name of the test
		rink          Rink   // the rink we want to test
		expectedValue string // gthe string value we expect
	}{
		{"Bairel", RinkBairel, "BAIREL"},
		{"UPMC", RinkUPMC, "UPMC"},
	}

	for _, testData := range tests {
		testFramework.Run(testData.testName, func(testFramework *testing.T) {
			receivedString := testData.rink.String()

			if receivedString != testData.expectedValue {
				testFramework.Errorf("Expected %v, got %v", testData.expectedValue, receivedString)
			}
		})
	}
}

func TestAllRinks(testFramework *testing.T) {

	rinks := AllRinks()

	// correct number
	if len(rinks) != 2 {
		testFramework.Errorf("Expected 2 rinks, got %d", len(rinks))
	}

	// correct contents
	if !slices.Contains(rinks, RinkBairel) {
		testFramework.Error("AllRinks does not contain Bairel")
	}

	if !slices.Contains(rinks, RinkUPMC) {
		testFramework.Error("AllRinks does not contain UPMC")
	}
}
