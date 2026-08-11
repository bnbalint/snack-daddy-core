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

func (mock *mockDB) GetAllUsers(ctx context.Context) ([]models.User, error) {
	if mock.getAllUsersFunc != nil {
		return mock.getAllUsersFunc(ctx)
	}
	return nil, nil
}

func (mock *mockDB) GetUserById(ctx context.Context, userId int) (*models.User, error) {
	if mock.getUserByIdFunc != nil {
		return mock.getUserByIdFunc(ctx, userId)
	}
	return nil, nil
}

func (mock *mockDB) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	if mock.addUserFunc != nil {
		return mock.addUserFunc(ctx, user)
	}
	return nil, nil
}

func (mock *mockDB) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if mock.updateUserFunc != nil {
		return mock.updateUserFunc(ctx, user)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllUsers
// .
// Tests
//   - success
//   - database error
func Test_GetAllUsers(testFramework *testing.T) {

	TEAM_MULES := models.Team{
		ID:             1,
		Name:           "Mules",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
	}

	// Define the tests
	tests := []struct {
		name           string
		mockUsers      []models.User
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain users
	}{
		{
			name: "success",
			mockUsers: []models.User{
				{ID: 1, FirstName: "Roger", LastName: "Hogwarts", Email: "r.h@gmail.com", Teams: []models.Team{TEAM_MULES}, Allergies: []models.Ingredient{}},
				{ID: 2, FirstName: "Brandi", LastName: "Hogwarta", Email: "b@gmail.com", Teams: []models.Team{}, Allergies: []models.Ingredient{}},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockUsers:      nil,
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
				getAllUsersFunc: func(ctx context.Context) ([]models.User, error) {
					return testData.mockUsers, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/users", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllUsers(ctx)
			if err != nil {
				testFramework.Errorf("GetAllUsers returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var users []models.User
				if err := json.Unmarshal(rec.Body.Bytes(), &users); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(users) != len(testData.mockUsers) {
					testFramework.Errorf("expected %d users, got %d", len(testData.mockUsers), len(users))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddUser
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func Test_AddUser(testFramework *testing.T) {

	TEAM_MULES := models.Team{
		ID:             1,
		Name:           "Mules",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
	}

	ALLERGY_PECAN := models.Ingredient{
		ID:   1,
		Name: "Pecan",
	}

	// Define the tests
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		mockError      error
		expectBody     bool
		mockReturnUser *models.User
	}{
		{
			name:           "success",
			requestBody:    `{"FirstName": "Roger", "LastName": "Hogwarts", "Email": "r.h@gmail.com", "Teams": [{"Name":"Mules", "Rink":"BAIREL", "Level":"D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": [{"Name": "Pecan"}]}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnUser: &models.User{
				ID:        1,
				FirstName: "Roger",
				LastName:  "Hogwarts",
				Email:     "r.h@gmail.com",
				Teams:     []models.Team{TEAM_MULES},
				Allergies: []models.Ingredient{ALLERGY_PECAN},
			},
		},
		{
			name:           "success_noAllergies",
			requestBody:    `{"FirstName": "Roger", "LastName": "Hogwarts", "Email": "r.h@gmail.com", "Teams": [{"Name":"Mules", "Rink":"BAIREL", "Level":"D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": []}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnUser: &models.User{
				ID:        1,
				FirstName: "Roger",
				LastName:  "Hogwarts",
				Email:     "r.h@gmail.com",
				Teams:     []models.Team{TEAM_MULES},
				Allergies: []models.Ingredient{},
			},
		},
		{
			name:           "success_noTeams",
			requestBody:    `{"FirstName": "Roger", "LastName": "Hogwarts", "Email": "r.h@gmail.com", "Teams": [], "Allergies": [{"Name": "Pecan"}]}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnUser: &models.User{
				ID:        1,
				FirstName: "Roger",
				LastName:  "Hogwarts",
				Email:     "r.h@gmail.com",
				Teams:     []models.Team{},
				Allergies: []models.Ingredient{ALLERGY_PECAN},
			},
		},
		{
			name:           "success_noTeams_noAllergies",
			requestBody:    `{"FirstName": "Roger", "LastName": "Hogwarts", "Email": "r.h@gmail.com"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnUser: &models.User{
				ID:        1,
				FirstName: "Roger",
				LastName:  "Hogwarts",
				Email:     "r.h@gmail.com",
				Teams:     []models.Team{},
				Allergies: []models.Ingredient{},
			},
		},
		{
			name:           "bind error",
			requestBody:    "invalid json",
			expectedStatus: http.StatusUnsupportedMediaType,
			mockError:      nil,
			expectBody:     false,
			mockReturnUser: nil,
		},
		{
			name:           "conflict error",
			requestBody:    `{"FirstName": "Roger", "LastName": "Hogwarts", "Email": "r.h@gmail.com", "Teams": [{"Name": "Mules", "Rink":"BAIREL", "Level": "D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": []}`,
			expectedStatus: http.StatusConflict,
			mockError:      &database_errors.ConflictError{},
			expectBody:     false,
			mockReturnUser: nil,
		},
		{
			name:           "database error",
			requestBody:    `{"FirstName": "Roger", "LastName": "Hogwarts", "Email": "r.h@gmail.com", "Teams": [{"Name": "Mules", "Rink": "BAIREL", "Level": "D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": []}`,
			expectedStatus: http.StatusInternalServerError,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:     false,
			mockReturnUser: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addUserFunc: func(ctx context.Context, user *models.User) (*models.User, error) {
					return testData.mockReturnUser, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddUser(ctx)
			if err != nil {
				testFramework.Errorf("AddUser returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var user models.User
				if err := json.Unmarshal(rec.Body.Bytes(), &user); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if user.ID != testData.mockReturnUser.ID || user.FirstName != testData.mockReturnUser.FirstName {
					testFramework.Errorf("expected user %+v, got %+v", testData.mockReturnUser, user)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// GetUserById
// .
// Tests
//   - success
//   - userId is not an integer
//   - userId is less than or equal to 0
//   - user not found
//   - generic database error
func Test_GetUserById(testFramework *testing.T) {

	TEAM_MULES := models.Team{
		ID:             1,
		Name:           "Mules",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
	}

	// Define the tests
	tests := []struct {
		name           string
		userId         string
		expectedStatus int
		mockError      error
		expectedUser   *models.User // true if body should contain users
	}{
		{
			name:           "success",
			userId:         "1",
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedUser:   &models.User{ID: 1, FirstName: "Roger", LastName: "Hogwarts", Email: "r.h@gmail.com", Teams: []models.Team{TEAM_MULES}, Allergies: []models.Ingredient{}},
		},
		{
			name:           "userId_not_valid_integer",
			userId:         "bob",
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
			expectedUser:   nil,
		},
		{
			name:           "userId_less_than_0",
			userId:         "-2",
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
			expectedUser:   nil,
		},
		{
			name:           "user_not_found",
			userId:         "500",
			expectedStatus: http.StatusNotFound,
			mockError:      &database_errors.NotFoundError{},
			expectedUser:   nil,
		},
		{
			name:           "databse error",
			userId:         "1",
			expectedStatus: http.StatusInternalServerError,
			mockError:      echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectedUser:   nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				getUserByIdFunc: func(ctx context.Context, userId int) (*models.User, error) {
					return testData.expectedUser, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)
			ctx.SetPath("users/:userId")
			ctx.SetParamNames("userId")
			ctx.SetParamValues(testData.userId)

			// Call handler
			err := server.GetUserById(ctx)
			if err != nil {
				testFramework.Errorf("GetAllUsers returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedUser != nil {
				var user models.User
				if err := json.Unmarshal(rec.Body.Bytes(), &user); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if user.ID != testData.expectedUser.ID {
					testFramework.Errorf("expected user to have ID %v, got %d", testData.expectedUser.ID, user.ID)
				}
			}
		})
	}
}
