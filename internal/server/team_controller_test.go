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

func (m *mockDB) Ready() bool {
	return true
}

func (m *mockDB) GetAllTeams(ctx context.Context) ([]models.Team, error) {
	if m.getAllTeamsFunc != nil {
		return m.getAllTeamsFunc(ctx)
	}
	return nil, nil
}

func (m *mockDB) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	if m.addTeamFunc != nil {
		return m.addTeamFunc(ctx, team)
	}
	return nil, nil
}

// Stub implementations for other methods
func (m *mockDB) GetAllUsers(ctx context.Context) ([]models.User, error) {
	panic("not implemented")
}

func (m *mockDB) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("not implemented")
}

func (m *mockDB) GetAllSnacks(ctx context.Context) ([]models.Snack, error) {
	panic("not implemented")
}

func (m *mockDB) AddSnack(ctx context.Context, snack *models.Snack) (*models.Snack, error) {
	panic("not implemented")
}

func (m *mockDB) GetAllAllergies(ctx context.Context) ([]models.Allergy, error) {
	panic("not implemented")
}

func (m *mockDB) AddAllergy(ctx context.Context, allergy *models.Allergy) (*models.Allergy, error) {
	panic("not implemented")
}

func TestGetAllTeams(t *testing.T) {
	tests := []struct {
		name           string
		mockTeams      []models.Team
		mockError      error
		expectedStatus int
		expectedBody   bool // true if body should contain teams
	}{
		{
			name: "success",
			mockTeams: []models.Team{
				{ID: 1, Name: "Team A"},
				{ID: 2, Name: "Team B"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockTeams:      nil,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mock := &mockDB{
				getAllTeamsFunc: func(ctx context.Context) ([]models.Team, error) {
					return tt.mockTeams, tt.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/teams", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)

			// Call handler
			err := server.GetAllTeams(ctx)
			if err != nil {
				t.Errorf("GetAllTeams returned error: %v", err)
			}

			// Check status
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Check body if expected
			if tt.expectedBody {
				var teams []models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &teams); err != nil {
					t.Errorf("failed to unmarshal response: %v", err)
				}
				if len(teams) != len(tt.mockTeams) {
					t.Errorf("expected %d teams, got %d", len(tt.mockTeams), len(teams))
				}
			}
		})
	}
}

func TestAddTeam(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		mockReturnTeam *models.Team
		mockError      error
		expectedStatus int
		expectBody     bool
	}{
		{
			name:        "success",
			requestBody: `{"name":"New Team","rink":"Rink A","level":"Beginner"}`,
			mockReturnTeam: &models.Team{
				ID:    1,
				Name:  "New Team",
				Rink:  "Rink A",
				Level: "Beginner",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectBody:     true,
		},
		{
			name:           "bind error",
			requestBody:    "invalid json",
			mockReturnTeam: nil,
			mockError:      nil,
			expectedStatus: http.StatusUnsupportedMediaType,
			expectBody:     false,
		},
		{
			name:           "conflict error",
			requestBody:    `{"name":"Existing Team","rink":"Rink B","level":"Advanced"}`,
			mockReturnTeam: nil,
			mockError:      &database_errors.ConflictError{},
			expectedStatus: http.StatusConflict,
			expectBody:     false,
		},
		{
			name:           "database error",
			requestBody:    `{"name":"Team C","rink":"Rink C","level":"Intermediate"}`,
			mockReturnTeam: nil,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedStatus: http.StatusInternalServerError,
			expectBody:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mock := &mockDB{
				addTeamFunc: func(ctx context.Context, team *models.Team) (*models.Team, error) {
					return tt.mockReturnTeam, tt.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request body
			body := []byte(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)

			// Call handler
			err := server.AddTeam(ctx)
			if err != nil {
				t.Errorf("AddTeam returned error: %v", err)
			}

			// Check status
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Check body if expected
			if tt.expectBody {
				var team models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &team); err != nil {
					t.Errorf("failed to unmarshal response: %v", err)
				}
				if team.ID != tt.mockReturnTeam.ID || team.Name != tt.mockReturnTeam.Name {
					t.Errorf("expected team %+v, got %+v", tt.mockReturnTeam, team)
				}
			}
		})
	}
}
