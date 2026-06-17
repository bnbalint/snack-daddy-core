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
func (mock *mockDB) GetAllIngredients(ctx context.Context) ([]models.Ingredient, error) {
	if mock.getAllIngredientsFunc != nil {
		return mock.getAllIngredientsFunc(ctx)
	}
	return nil, nil
}

/**
* Call the testing function below that contains all of the test scenarios
 */
func (mock *mockDB) AddIngredient(ctx context.Context, ingredient *models.Ingredient) (*models.Ingredient, error) {
	if mock.addIngredientsFunc != nil {
		return mock.addIngredientsFunc(ctx, ingredient)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllIngredients
// .
// Tests
//   - success
//   - database error
func Test_GetAllIngredients(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name            string
		mockIngredients []models.Ingredient
		expectedStatus  int
		mockError       error
		expectedBody    bool // true if body should contain ingredients
	}{
		{
			name: "success",
			mockIngredients: []models.Ingredient{
				{ID: 1, Name: "Almond"},
				{ID: 2, Name: "Wheat"},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:            "database error",
			mockIngredients: nil,
			expectedStatus:  http.StatusInternalServerError,
			mockError:       echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedBody:    false,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				getAllIngredientsFunc: func(ctx context.Context) ([]models.Ingredient, error) {
					return testData.mockIngredients, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/ingredients", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllIngredients(ctx)
			if err != nil {
				testFramework.Errorf("GetAllIngredients returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var ingredients []models.Ingredient
				if err := json.Unmarshal(rec.Body.Bytes(), &ingredients); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(ingredients) != len(testData.mockIngredients) {
					testFramework.Errorf("expected %d ingredients, got %d", len(testData.mockIngredients), len(ingredients))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddIngredient
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func Test_AddIngredient(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name                 string
		requestBody          string
		expectedStatus       int
		mockError            error
		expectBody           bool
		mockReturnIngredient *models.Ingredient
	}{
		{
			name:           "success",
			requestBody:    `{"Name": "Peanut"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnIngredient: &models.Ingredient{
				ID:   1,
				Name: "Peanut",
			},
		},
		{
			name:                 "bind error",
			requestBody:          "invalid json",
			expectedStatus:       http.StatusUnsupportedMediaType,
			mockError:            nil,
			expectBody:           false,
			mockReturnIngredient: nil,
		},
		{
			name:                 "conflict error",
			requestBody:          `{"Name": "Peanut"}`,
			expectedStatus:       http.StatusConflict,
			mockError:            &database_errors.ConflictError{},
			expectBody:           false,
			mockReturnIngredient: nil,
		},
		{
			name:                 "database error",
			requestBody:          `{"Name": "Peanut"}`,
			expectedStatus:       http.StatusInternalServerError,
			mockError:            echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:           false,
			mockReturnIngredient: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addIngredientsFunc: func(ctx context.Context, ingredient *models.Ingredient) (*models.Ingredient, error) {
					return testData.mockReturnIngredient, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/ingredients", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddIngredient(ctx)
			if err != nil {
				testFramework.Errorf("AddIngredient returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var ingredient models.Ingredient
				if err := json.Unmarshal(rec.Body.Bytes(), &ingredient); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if ingredient.ID != testData.mockReturnIngredient.ID || ingredient.Name != testData.mockReturnIngredient.Name {
					testFramework.Errorf("expected ingredient %+v, got %+v", testData.mockReturnIngredient, ingredient)
				}
			}
		})
	}
}
