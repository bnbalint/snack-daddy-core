package server

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	database_errors "snack-daddy-core/internal/database/errors"
	"snack-daddy-core/internal/models"

	"github.com/labstack/echo/v4"
)

func (mock *mockDB) GetAllSnacks(ctx context.Context) ([]models.Snack, error) {
	if mock.getAllSnacksFunc != nil {
		return mock.getAllSnacksFunc(ctx)
	}
	return nil, nil
}

func (mock *mockDB) AddSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error) {
	if mock.addSnackFunc != nil {
		return mock.addSnackFunc(ctx, snack)
	}
	return nil, nil
}

func (mock *mockDB) UpdateSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error) {
	if mock.updateSnackFunc != nil {
		return mock.updateSnackFunc(ctx, snack)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllSnacks
// .
// Tests
//   - success
//   - database error
func Test_GetAllSnacks(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		mockSnacks     []models.Snack
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain snacks
	}{
		{
			name: "success",
			mockSnacks: []models.Snack{
				{ID: 1, Name: "Rice Crispie Treat", Sweet: true, Savory: false, Difficulty: 2, RecipeUrl: ""},
				{ID: 2, Name: "Bacon Crackers", Sweet: false, Savory: true, Difficulty: 3, RecipeUrl: ""},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockSnacks:     nil,
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
				getAllSnacksFunc: func(ctx context.Context) ([]models.Snack, error) {
					return testData.mockSnacks, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/snacks", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllSnacks(ctx)
			if err != nil {
				testFramework.Errorf("GetAllSnacks returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var snacks []models.Snack
				if err := json.Unmarshal(rec.Body.Bytes(), &snacks); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(snacks) != len(testData.mockSnacks) {
					testFramework.Errorf("expected %d snacks, got %d", len(testData.mockSnacks), len(snacks))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddSnack
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func Test_AddSnack(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name            string
		requestBody     string
		expectedStatus  int
		mockError       error
		expectBody      bool
		mockReturnSnack *models.Snack
	}{
		{
			name:           "success",
			requestBody:    `{"Name": "Rice Crispie Treat", "Sweet": true, "Savory": false, "Difficulty": 2}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnSnack: &models.Snack{
				ID:         1,
				Name:       "Rice Crispie Treat",
				Sweet:      true,
				Savory:     false,
				Difficulty: 2,
			},
		},
		{
			name:            "bind error",
			requestBody:     "invalid json",
			expectedStatus:  http.StatusUnsupportedMediaType,
			mockError:       nil,
			expectBody:      false,
			mockReturnSnack: nil,
		},
		{
			name:            "conflict error",
			requestBody:     `{"Name": "Rice Crispie Treat", "Sweet": true, "Savory": false, "Difficulty": 2}`,
			expectedStatus:  http.StatusConflict,
			mockError:       &database_errors.ConflictError{},
			expectBody:      false,
			mockReturnSnack: nil,
		},
		{
			name:            "database error",
			requestBody:     `{"Name": "Rice Crispie Treat", "Sweet": true, "Savory": false, "Difficulty": 2}`,
			expectedStatus:  http.StatusInternalServerError,
			mockError:       echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:      false,
			mockReturnSnack: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addSnackFunc: func(ctx context.Context, snack *models.Snack) (*models.Snack, error) {
					return testData.mockReturnSnack, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/snacks", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddSnack(ctx)
			if err != nil {
				testFramework.Errorf("AddSnack returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var snack models.Snack
				if err := json.Unmarshal(rec.Body.Bytes(), &snack); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if snack.ID != testData.mockReturnSnack.ID || snack.Name != testData.mockReturnSnack.Name {
					testFramework.Errorf("expected snack %+v, got %+v", testData.mockReturnSnack, snack)
				}
			}
		})
	}
}
