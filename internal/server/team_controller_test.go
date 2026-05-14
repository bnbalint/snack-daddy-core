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

// mockDB implements SnackDaddyDatabaseClient for testing
type mockDB struct {
	getAllTeamsFunc func(ctx context.Context) ([]models.Team, error)
	addTeamFunc     func(ctx context.Context, team *models.Team) (*models.Team, error)
}

func (mock *mockDB) Ready() bool {
	return true
}

func (mock *mockDB) GetAllTeams(ctx context.Context) ([]models.Team, error) {
	if mock.getAllTeamsFunc != nil {
		return mock.getAllTeamsFunc(ctx)
	}
	return nil, nil
}

func (mock *mockDB) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	if mock.addTeamFunc != nil {
		return mock.addTeamFunc(ctx, team)
	}
	return nil, nil
}

// Stub implementations for other methods
func (mock *mockDB) GetAllUsers(ctx context.Context) ([]models.User, error) {
	panic("not implemented")
}

func (mock *mockDB) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("not implemented")
}

func (mock *mockDB) GetAllSnacks(ctx context.Context) ([]models.Snack, error) {
	panic("not implemented")
}

func (mock *mockDB) AddSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error) {
	panic("not implemented")
}

func (mock *mockDB) GetAllAllergies(ctx context.Context) ([]models.Allergy, error) {
	panic("not implemented")
}

func (mock *mockDB) AddAllergy(ctx context.Context, allergy *models.Allergy) (*models.Allergy, error) {
	panic("not implemented")
}

// ---------------------------------------------------------------------
// GetAllTeams
// .
// Tests
//   - success
//   - database error
func TestGetAllTeams(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		mockTeams      []models.Team
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain teams
	}{
		{
			name: "success",
			mockTeams: []models.Team{
				{ID: 1, Name: "Team A"},
				{ID: 2, Name: "Team B"},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockTeams:      nil,
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
				getAllTeamsFunc: func(ctx context.Context) ([]models.Team, error) {
					return testData.mockTeams, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/teams", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllTeams(ctx)
			if err != nil {
				testFramework.Errorf("GetAllTeams returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var teams []models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &teams); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(teams) != len(testData.mockTeams) {
					testFramework.Errorf("expected %d teams, got %d", len(testData.mockTeams), len(teams))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddTeam
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func TestAddTeam(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		mockError      error
		expectBody     bool
		mockReturnTeam *models.Team
	}{
		{
			name:           "success",
			requestBody:    `{"name":"Mules","rink":"Baierl","level":"D5"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnTeam: &models.Team{
				ID:    1,
				Name:  "Mules",
				Rink:  "Baierl",
				Level: "D5",
			},
		},
		{
			name:           "bind error",
			requestBody:    "invalid json",
			expectedStatus: http.StatusUnsupportedMediaType,
			mockError:      nil,
			expectBody:     false,
			mockReturnTeam: nil,
		},
		{
			name:           "conflict error",
			requestBody:    `{"name":"Mules","rink":"Baierl","level":"D5"}`,
			expectedStatus: http.StatusConflict,
			mockError:      &database_errors.ConflictError{},
			expectBody:     false,
			mockReturnTeam: nil,
		},
		{
			name:           "database error",
			requestBody:    `{"name":"Mules","rink":"Baierl","level":"D5"}`,
			expectedStatus: http.StatusInternalServerError,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:     false,
			mockReturnTeam: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addTeamFunc: func(ctx context.Context, team *models.Team) (*models.Team, error) {
					return testData.mockReturnTeam, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddTeam(ctx)
			if err != nil {
				testFramework.Errorf("AddTeam returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var team models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &team); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if team.ID != testData.mockReturnTeam.ID || team.Name != testData.mockReturnTeam.Name {
					testFramework.Errorf("expected team %+v, got %+v", testData.mockReturnTeam, team)
				}
			}
		})
	}
}
