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
func (mock *mockDB) GetAllAllergies(ctx context.Context) ([]models.Allergy, error) {
	if mock.getAllAllergiesFunc != nil {
		return mock.getAllAllergiesFunc(ctx)
	}
	return nil, nil
}

/**
* Call the testing function below that contains all of the test scenarios
 */
func (mock *mockDB) AddAllergy(ctx context.Context, allergy *models.Allergy) (*models.Allergy, error) {
	if mock.addAllergyFunc != nil {
		return mock.addAllergyFunc(ctx, allergy)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllAllergies
// .
// Tests
//   - success
//   - database error
func TestGetAllAllergies(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		mockAllergies  []models.Allergy
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain allergies
	}{
		{
			name: "success",
			mockAllergies: []models.Allergy{
				{ID: 1, Name: "Almond"},
				{ID: 2, Name: "Wheat"},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockAllergies:  nil,
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
				getAllAllergiesFunc: func(ctx context.Context) ([]models.Allergy, error) {
					return testData.mockAllergies, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/allergies", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllAllergies(ctx)
			if err != nil {
				testFramework.Errorf("GetAllAllergies returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var allergies []models.Allergy
				if err := json.Unmarshal(rec.Body.Bytes(), &allergies); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(allergies) != len(testData.mockAllergies) {
					testFramework.Errorf("expected %d allergies, got %d", len(testData.mockAllergies), len(allergies))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddAllergy
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func TestAddAllergy(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name              string
		requestBody       string
		expectedStatus    int
		mockError         error
		expectBody        bool
		mockReturnAllergy *models.Allergy
	}{
		{
			name:           "success",
			requestBody:    `{"Name": "Peanut"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnAllergy: &models.Allergy{
				ID:   1,
				Name: "Peanut",
			},
		},
		{
			name:              "bind error",
			requestBody:       "invalid json",
			expectedStatus:    http.StatusUnsupportedMediaType,
			mockError:         nil,
			expectBody:        false,
			mockReturnAllergy: nil,
		},
		{
			name:              "conflict error",
			requestBody:       `{"Name": "Peanut"}`,
			expectedStatus:    http.StatusConflict,
			mockError:         &database_errors.ConflictError{},
			expectBody:        false,
			mockReturnAllergy: nil,
		},
		{
			name:              "database error",
			requestBody:       `{"Name": "Peanut"}`,
			expectedStatus:    http.StatusInternalServerError,
			mockError:         echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:        false,
			mockReturnAllergy: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addAllergyFunc: func(ctx context.Context, allergy *models.Allergy) (*models.Allergy, error) {
					return testData.mockReturnAllergy, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/allergies", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddAllergy(ctx)
			if err != nil {
				testFramework.Errorf("AddAllergy returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var allergy models.Allergy
				if err := json.Unmarshal(rec.Body.Bytes(), &allergy); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if allergy.ID != testData.mockReturnAllergy.ID || allergy.Name != testData.mockReturnAllergy.Name {
					testFramework.Errorf("expected allergy %+v, got %+v", testData.mockReturnAllergy, allergy)
				}
			}
		})
	}
}
