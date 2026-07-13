package server

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

/**
* Call the testing function below that contains all of the test scenarios
 */
func (mock *mockDB) GetSnackLog(ctx context.Context) ([]models.SnackLog, error) {
	if mock.getSnackLogFunc != nil {
		return mock.getSnackLogFunc(ctx)
	}
	return nil, nil
}

/**
* Call the testing function below that contains all of the test scenarios
 */
func (mock *mockDB) AddToSnackLog(ctx context.Context, snackLogEntry *models.SnackLog) (*models.SnackLog, error) {
	if mock.addToSnackLogFunc != nil {
		return mock.addToSnackLogFunc(ctx, snackLogEntry)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetSnackLog
// .
// Tests
//   - success
//   - database error
func Test_GetSnackLog(testFramework *testing.T) {

	SNACK := models.Snack{
		Name:       "Rice Crispie Treat",
		Sweet:      true,
		Savory:     false,
		Difficulty: 2,
		RecipeUrl:  "",
	}

	TEAM := models.Team{
		Name:           "Mules",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
	}

	DATE := time.Date(2026, time.June, 23, 0, 0, 0, 0, time.UTC)

	// Define the tests
	tests := []struct {
		name           string
		mockSnackLog   []models.SnackLog
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain snackLog entries
	}{
		{
			name: "success",
			mockSnackLog: []models.SnackLog{
				{ID: 1, SnackID: 1, Snack: SNACK, TeamID: 1, Team: TEAM, DateMade: DATE},
				{ID: 2, SnackID: 2, Snack: SNACK, TeamID: 1, Team: TEAM, DateMade: DATE},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockSnackLog:   nil,
			expectedStatus: http.StatusInternalServerError,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedBody:   false,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				getSnackLogFunc: func(ctx context.Context) ([]models.SnackLog, error) {
					return testData.mockSnackLog, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/snack-log", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetSnackLog(ctx)
			if err != nil {
				testFramework.Errorf("GetSnackLog returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var snackLog []models.SnackLog
				if err := json.Unmarshal(rec.Body.Bytes(), &snackLog); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(snackLog) != len(testData.mockSnackLog) {
					testFramework.Errorf("expected %d snack log entries, got %d", len(testData.mockSnackLog), len(snackLog))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddToSnackLog
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func Test_AddToSnackLog(testFramework *testing.T) {

	DATE := time.Date(2026, time.June, 23, 0, 0, 0, 0, time.UTC)

	// Define the tests
	tests := []struct {
		name            string
		requestBody     string
		expectedStatus  int
		mockError       error
		expectBody      bool
		mockReturnEntry *models.SnackLog
	}{
		{
			name:           "success",
			requestBody:    `{"ID": 1, "SnackID": 1, "TeamID": 1, "DateMade": "2026-06-23T00:00:00Z"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnEntry: &models.SnackLog{
				ID:       1,
				SnackID:  1,
				TeamID:   1,
				DateMade: DATE,
			},
		},
		{
			name:            "bind error",
			requestBody:     "invalid json",
			expectedStatus:  http.StatusUnsupportedMediaType,
			mockError:       nil,
			expectBody:      false,
			mockReturnEntry: nil,
		},
		{
			name:            "conflict error",
			requestBody:     `{"ID": 1, "SnackID": 1, "TeamID": 1, "DateMade": "2026-06-23T00:00:00Z"}`,
			expectedStatus:  http.StatusConflict,
			mockError:       &database_errors.ConflictError{},
			expectBody:      false,
			mockReturnEntry: nil,
		},
		{
			name:            "database error",
			requestBody:     `{"ID": 1, "SnackID": 1, "TeamID": 1, "DateMade": "2026-06-23T00:00:00Z"}`,
			expectedStatus:  http.StatusInternalServerError,
			mockError:       echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:      false,
			mockReturnEntry: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addToSnackLogFunc: func(ctx context.Context, entry *models.SnackLog) (*models.SnackLog, error) {
					return testData.mockReturnEntry, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request body
			body := []byte(testData.requestBody)
			request := httptest.NewRequest(http.MethodPost, "/snack-log", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddToSnackLog(ctx)
			if err != nil {
				testFramework.Errorf("AddToSnackLog returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var entry models.SnackLog
				if err := json.Unmarshal(rec.Body.Bytes(), &entry); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if entry.ID != testData.mockReturnEntry.ID {
					testFramework.Errorf("expected entry %+v, got %+v", testData.mockReturnEntry, entry)
				}
			}
		})
	}
}
