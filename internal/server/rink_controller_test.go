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
// GetAllRinks
// .
// Tests
//   - success
func Test_GetAllRinks(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		mockRinks      []models.Rink
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain rinks
	}{
		{
			name: "success",
			mockRinks: []models.Rink{
				models.RinkBairel,
				models.RinkUPMC,
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
			request := httptest.NewRequest(http.MethodGet, "/rinks", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllRinks(ctx)
			if err != nil {
				testFramework.Errorf("GetAllRinks returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var rinks []models.Rink
				if err := json.Unmarshal(rec.Body.Bytes(), &rinks); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(rinks) != len(testData.mockRinks) {
					testFramework.Errorf("expected %d rinks, got %d", len(testData.mockRinks), len(rinks))
				}
			}
		})
	}
}
