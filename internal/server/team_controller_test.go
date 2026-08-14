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

func (mock *mockDB) GetAllTeams(ctx context.Context) ([]models.Team, error) {
	if mock.getAllTeamsFunc != nil {
		return mock.getAllTeamsFunc(ctx)
	}
	return nil, nil
}

func (mock *mockDB) GetTeamById(ctx context.Context, teamId int) (*models.Team, error) {
	if mock.getTeamByIdFunc != nil {
		return mock.getTeamByIdFunc(ctx, teamId)
	}
	return nil, nil
}

func (mock *mockDB) AddTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	if mock.addTeamFunc != nil {
		return mock.addTeamFunc(ctx, team)
	}
	return nil, nil
}

func (mock *mockDB) UpdateTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	if mock.updateTeamFunc != nil {
		return mock.updateTeamFunc(ctx, team)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllTeams
// .
// Tests
//   - success
//   - database error
func Test_GetAllTeams(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name           string
		expectedBody   []models.Team
		expectedStatus int
		mockError      error
	}{
		{
			name: "success",
			expectedBody: []models.Team{
				{ID: 1, Name: "Team A"},
				{ID: 2, Name: "Team B"},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
		},
		{
			name:           "database error",
			expectedBody:   nil,
			expectedStatus: http.StatusInternalServerError,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				getAllTeamsFunc: func(ctx context.Context) ([]models.Team, error) {
					return testData.expectedBody, testData.mockError
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
			if testData.expectedBody != nil {
				var teams []models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &teams); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(teams) != len(testData.expectedBody) {
					testFramework.Errorf("expected %d teams, got %d", len(testData.expectedBody), len(teams))
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
func Test_AddTeam(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name             string
		requestBody      string
		expectedStatus   int
		addError         error
		expectBody       bool
		expectedResponse *models.Team
	}{
		{
			name:           "success",
			requestBody:    `{"name":"Mules","rink":"BAIREL","level":"D5"}`,
			expectedStatus: http.StatusCreated,
			addError:       nil,
			expectedResponse: &models.Team{
				ID:    1,
				Name:  "Mules",
				Rink:  models.RinkBairel,
				Level: models.LevelD5,
			},
		},
		{
			name:             "bind error",
			requestBody:      "invalid json",
			expectedStatus:   http.StatusUnsupportedMediaType,
			addError:         nil,
			expectedResponse: nil,
		},
		{
			name:             "conflict error",
			requestBody:      `{"name":"Mules","rink":"BAIREL","level":"D5"}`,
			expectedStatus:   http.StatusConflict,
			addError:         &database_errors.ConflictError{},
			expectedResponse: nil,
		},
		{
			name:             "database error",
			requestBody:      `{"name":"Mules","rink":"BAIREL","level":"D5"}`,
			expectedStatus:   http.StatusInternalServerError,
			addError:         echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedResponse: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addTeamFunc: func(ctx context.Context, team *models.Team) (*models.Team, error) {
					return testData.expectedResponse, testData.addError
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
			if testData.expectedResponse != nil {
				var team models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &team); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if team.ID != testData.expectedResponse.ID || team.Name != testData.expectedResponse.Name {
					testFramework.Errorf("expected team %+v, got %+v", testData.expectedResponse, team)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// UpdateTeam
// .
// Tests
//   - success
//   - bind error
//   - team does not have an ID
//   - team does not already exist
//   - conflict error
//   - database error
func Test_UpdateTeam(testFramework *testing.T) {

	// Define the tests
	tests := []struct {
		name             string
		requestBody      string
		expectedStatus   int
		getTeamByIdError error
		updateError      error
		expectedResponse *models.Team
	}{
		{
			name:             "success",
			requestBody:      `{"id": 1, "name":"Mules","rink":"BAIREL","level":"D5"}`,
			expectedStatus:   http.StatusOK,
			getTeamByIdError: nil,
			updateError:      nil,
			expectedResponse: &models.Team{
				ID:    1,
				Name:  "Mules",
				Rink:  models.RinkBairel,
				Level: models.LevelD5,
			},
		},
		{
			name:             "bind error",
			requestBody:      "invalid json",
			expectedStatus:   http.StatusUnsupportedMediaType,
			getTeamByIdError: nil,
			updateError:      nil,
			expectedResponse: nil,
		},
		{
			name:             "Missing TeamID",
			requestBody:      `{"name":"Mules","rink":"BAIREL","level":"D5"}`,
			expectedStatus:   http.StatusBadRequest,
			getTeamByIdError: nil,
			updateError:      nil,
			expectedResponse: nil,
		},
		{
			name:             "team does not exist",
			requestBody:      `{"id":1, "name":"Mules", "rink":"BAIREL", "level":"D5"}`,
			expectedStatus:   http.StatusBadRequest,
			getTeamByIdError: &database_errors.NotFoundError{},
			updateError:      nil,
			expectedResponse: nil,
		},
		{
			name:             "conflict error",
			requestBody:      `{"id": 1, "name":"Mules", "rink":"BAIREL", "level":"D5"}`,
			expectedStatus:   http.StatusConflict,
			getTeamByIdError: nil,
			updateError:      &database_errors.ConflictError{},
			expectedResponse: nil,
		},
		{
			name:             "database error",
			requestBody:      `{"id": 1, "name":"Mules", "rink":"BAIREL", "level":"D5"}`,
			expectedStatus:   http.StatusInternalServerError,
			getTeamByIdError: nil,
			updateError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedResponse: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				updateTeamFunc: func(ctx context.Context, team *models.Team) (*models.Team, error) {
					return testData.expectedResponse, testData.updateError
				},

				// we don't use the team returned from this, just the error
				getTeamByIdFunc: func(ctx context.Context, teamId int) (*models.Team, error) {
					return nil, testData.getTeamByIdError
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
			request := httptest.NewRequest(http.MethodPut, "/teams", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.UpdateTeam(ctx)
			if err != nil {
				testFramework.Errorf("UpdateTeam returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedResponse != nil {
				var team models.Team
				if err := json.Unmarshal(rec.Body.Bytes(), &team); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if team.ID != testData.expectedResponse.ID || team.Name != testData.expectedResponse.Name {
					testFramework.Errorf("expected team %+v, got %+v", testData.expectedResponse, team)
				}
			}
		})
	}
}
