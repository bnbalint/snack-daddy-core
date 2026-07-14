package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

// ---------------------------------------------------------------------
// GetAllLevels
// .
// Tests
//   - success
func Test_GetAllLevels(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		mockLevels     []models.Level
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain rinks
	}{
		{
			name: "success",
			mockLevels: []models.Level{
				models.LevelD5,
				models.LevelD4,
				models.LevelD3,
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     &mockDB{},
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/levels", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllLevels(ctx)
			if err != nil {
				testFramework.Errorf("GetAllLevels returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var levels []models.Level
				if err := json.Unmarshal(rec.Body.Bytes(), &levels); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(levels) != len(testData.mockLevels) {
					testFramework.Errorf("expected %d levels, got %d", len(testData.mockLevels), len(levels))
				}
			}
		})
	}
}
