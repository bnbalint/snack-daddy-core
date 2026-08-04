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

/**
* Call the testing function below that contains all of the test scenarios
 */
func (mock *mockDB) GetAllSuggestedAllergies(ctx context.Context) ([]models.SuggestedAllergy, error) {
	if mock.getAllSuggestedAllergiesFunc != nil {
		return mock.getAllSuggestedAllergiesFunc(ctx)
	}
	return nil, nil
}

/**
* Call the testing function below that contains all of the test scenarios
 */
func (mock *mockDB) AddSuggestedAllergy(ctx context.Context, suggestedAllergy *models.SuggestedAllergy) (*models.SuggestedAllergy, error) {
	if mock.addSuggestedAllergyFunc != nil {
		return mock.addSuggestedAllergyFunc(ctx, suggestedAllergy)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllSuggestedAllergies
// .
// Tests
//   - success
//   - database error
func Test_GetAllSuggestedAllergies(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name                   string
		mockSuggestedAllergies []models.SuggestedAllergy
		expectedStatus         int
		mockError              error
		expectedBody           bool // true if body should contain suggestedAllergies
	}{
		{
			name: "success",
			mockSuggestedAllergies: []models.SuggestedAllergy{
				{ID: 1, Name: "Pine nut"},
				{ID: 2, Name: "Gluten"},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:                   "database error",
			mockSuggestedAllergies: nil,
			expectedStatus:         http.StatusInternalServerError,
			mockError:              echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedBody:           false,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				getAllSuggestedAllergiesFunc: func(ctx context.Context) ([]models.SuggestedAllergy, error) {
					return testData.mockSuggestedAllergies, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/suggested-allergies", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllSuggestedAllergies(ctx)
			if err != nil {
				testFramework.Errorf("GetAllSuggestedAllergies returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var suggestedAllergies []models.SuggestedAllergy
				if err := json.Unmarshal(rec.Body.Bytes(), &suggestedAllergies); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(suggestedAllergies) != len(testData.mockSuggestedAllergies) {
					testFramework.Errorf("expected %d suggestedAllergies, got %d", len(testData.mockSuggestedAllergies), len(suggestedAllergies))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddSuggestedAllergy
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func Test_AddSuggestedAllergy(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name                 string
		requestBody          string
		expectedStatus       int
		mockError            error
		expectBody           bool
		mockReturnSuggestion *models.SuggestedAllergy
	}{
		{
			name:           "success",
			requestBody:    `{"Name": "Pine nut"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnSuggestion: &models.SuggestedAllergy{
				ID:   1,
				Name: "Pine nut",
			},
		},
		{
			name:                 "bind error",
			requestBody:          "invalid json",
			expectedStatus:       http.StatusUnsupportedMediaType,
			mockError:            nil,
			expectBody:           false,
			mockReturnSuggestion: nil,
		},
		{
			name:                 "conflict error",
			requestBody:          `{"Name": "Pine nut"}`,
			expectedStatus:       http.StatusConflict,
			mockError:            &database_errors.ConflictError{},
			expectBody:           false,
			mockReturnSuggestion: nil,
		},
		{
			name:                 "database error",
			requestBody:          `{"Name": "Pine nut"}`,
			expectedStatus:       http.StatusInternalServerError,
			mockError:            echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:           false,
			mockReturnSuggestion: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addSuggestedAllergyFunc: func(ctx context.Context, suggestion *models.SuggestedAllergy) (*models.SuggestedAllergy, error) {
					return testData.mockReturnSuggestion, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/suggested-allergies", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddSuggestedAllergy(ctx)
			if err != nil {
				testFramework.Errorf("AddSuggestedAllergy returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var suggestion models.SuggestedAllergy
				if err := json.Unmarshal(rec.Body.Bytes(), &suggestion); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if suggestion.ID != testData.mockReturnSuggestion.ID || suggestion.Name != testData.mockReturnSuggestion.Name {
					testFramework.Errorf("expected suggestedAllergy %+v, got %+v", testData.mockReturnSuggestion, suggestion)
				}
			}
		})
	}
}
