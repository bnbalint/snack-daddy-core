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

func (mock *mockDB) GetAllUserSnackRankings(ctx context.Context) ([]models.UserSnackRanking, error) {
	if mock.getAllUserSnackRankingsFunc != nil {
		return mock.getAllUserSnackRankingsFunc(ctx)
	}
	return nil, nil
}

func (mock *mockDB) GetUserSnackRankingsByUserId(ctx context.Context, userId int) ([]models.UserSnackRanking, error) {
	if mock.getUserSnackRankingsByUserIdFunc != nil {
		return mock.getUserSnackRankingsByUserIdFunc(ctx, userId)
	}
	return nil, nil
}

func (mock *mockDB) AddUserSnackRanking(ctx context.Context, user *models.UserSnackRanking) (*models.UserSnackRanking, error) {
	if mock.addUserSnackRankingFunc != nil {
		return mock.addUserSnackRankingFunc(ctx, user)
	}
	return nil, nil
}

// ---------------------------------------------------------------------
// GetAllUserSnackRankings
// .
// Tests
//   - success
//   - database error
func Test_GetAllUserSnackRankings(testFramework *testing.T) {

	//---------------------------------------------
	//  CONSTANTS - for creating the userSnackRanking
	//
	TIME, _ := time.Parse(time.RFC3339, "2026-07-07T00:00:00Z")

	ingredient1 := models.Ingredient{
		ID:        1,
		Name:      "Rice Crispy Cereal",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient2 := models.Ingredient{
		ID:        2,
		Name:      "Margarine",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient3 := models.Ingredient{
		ID:        3,
		Name:      "Marshmallow",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient4 := models.Ingredient{
		ID:        4,
		Name:      "Vanilla",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	SNACK := models.Snack{
		ID:          1,
		Name:        "Rice Crispie Treat",
		Sweet:       true,
		Savory:      false,
		Difficulty:  2,
		RecipeUrl:   "",
		Ingredients: []models.Ingredient{ingredient1, ingredient2, ingredient3, ingredient4},
		CreatedAt:   TIME,
		UpdatedAt:   TIME,
	}

	PECAN := models.Ingredient{
		ID:        1,
		Name:      "Pecan",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	TEAM_MULES := models.Team{
		ID:             1,
		Name:           "Mules",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
		CreatedAt:      TIME,
		UpdatedAt:      TIME,
	}

	USER := models.User{
		ID:        1,
		FirstName: "Roger",
		LastName:  "Hogwarts",
		Email:     "r.h@gmail.com",
		Teams:     []models.Team{TEAM_MULES},
		Allergies: []models.Ingredient{PECAN},
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	// Define the tests
	tests := []struct {
		name           string
		mockRankings   []models.UserSnackRanking
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain userSnackRankings
	}{
		{
			name: "success",
			mockRankings: []models.UserSnackRanking{
				{SnackID: 1, Snack: SNACK, UserID: 1, User: USER, Rank: models.SnackRank10, CreatedAt: TIME, UpdatedAt: TIME},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "database error",
			mockRankings:   nil,
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
				getAllUserSnackRankingsFunc: func(ctx context.Context) ([]models.UserSnackRanking, error) {
					return testData.mockRankings, testData.mockError
				},
			}

			// Create server
			logger := slog.New(slog.DiscardHandler)
			server := &SnackDaddyEchoServer{
				DB:     mock,
				Logger: logger,
			}

			// Create request
			request := httptest.NewRequest(http.MethodGet, "/snack-rankings", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.GetAllUserSnackRankings(ctx)
			if err != nil {
				testFramework.Errorf("GetAllUserSnackRankings returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var rankings []models.UserSnackRanking
				if err := json.Unmarshal(rec.Body.Bytes(), &rankings); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(rankings) != len(testData.mockRankings) {
					testFramework.Errorf("expected %d rankings, got %d", len(testData.mockRankings), len(rankings))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// AddUserSnackRanking
// .
// Tests
//   - success
//   - bind error
//   - conflict error
//   - database error
func Test_AddUserSnackRanking(testFramework *testing.T) {

	//---------------------------------------------
	//  CONSTANTS - for creating the userSnackRanking
	//

	SNACK := models.Snack{
		ID:         1,
		Name:       "Rice Crispie Treat",
		Sweet:      true,
		Savory:     false,
		Difficulty: 2,
		RecipeUrl:  "",
	}

	PECAN := models.Ingredient{
		ID:   1,
		Name: "Pecan",
	}

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

	USER := models.User{
		ID:        1,
		FirstName: "Roger",
		LastName:  "Hogwarts",
		Email:     "r.h@gmail.com",
		Teams:     []models.Team{TEAM_MULES},
		Allergies: []models.Ingredient{PECAN},
	}

	// Define the tests
	tests := []struct {
		name              string
		requestBody       string
		expectedStatus    int
		mockError         error
		expectBody        bool
		mockReturnRanking *models.UserSnackRanking
	}{
		{
			name:           "success",
			requestBody:    `{"snack_id": 1, "snack": {"ID": 1, "Name": "Rice Crispie Treat", "Sweet": true, "Savory": false, "Difficulty": 2}, "user_id": 1, "user": {"first_name_": "Roger", "last_name": "Hogwarts", "email": "r.h@gmail.com", "Teams": [{"Name":"Mules", "Rink":"BAIREL", "Level":"D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": [{"Name": "Pecan"}]}, "rank": "RANK_10"}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
			expectBody:     true,
			mockReturnRanking: &models.UserSnackRanking{
				SnackID: 1,
				Snack:   SNACK,
				UserID:  1,
				User:    USER,
				Rank:    models.SnackRank10,
			},
		},
		{
			name:              "bind error",
			requestBody:       "invalid json",
			expectedStatus:    http.StatusUnsupportedMediaType,
			mockError:         nil,
			expectBody:        false,
			mockReturnRanking: nil,
		},
		{
			name:              "conflict error",
			requestBody:       `{"snack_id": 1, "snack": {"ID": 1, "Name": "Rice Crispie Treat", "Sweet": true, "Savory": false, "Difficulty": 2}, "user_id": 1, "user": {"first_name_": "Roger", "last_name": "Hogwarts", "email": "r.h@gmail.com", "Teams": [{"Name":"Mules", "Rink":"BAIREL", "Level":"D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": [{"Name": "Pecan"}]}, "rank": "RANK_10"}`,
			expectedStatus:    http.StatusConflict,
			mockError:         &database_errors.ConflictError{},
			expectBody:        false,
			mockReturnRanking: nil,
		},
		{
			name:              "database error",
			requestBody:       `{"snack_id": 1, "snack": {"ID": 1, "Name": "Rice Crispie Treat", "Sweet": true, "Savory": false, "Difficulty": 2}, "user_id": 1, "user": {"first_name_": "Roger", "last_name": "Hogwarts", "email": "r.h@gmail.com", "Teams": [{"Name":"Mules", "Rink":"BAIREL", "Level":"D5", "PrimaryColor": "#b88907", "SecondaryColor": "#000000", "TernaryColor": "#c42323", "LogoUrl": ""}], "Allergies": [{"Name": "Pecan"}]}, "rank": "RANK_10"}`,
			expectedStatus:    http.StatusInternalServerError,
			mockError:         echo.NewHTTPError(http.StatusInternalServerError, "db error"),
			expectBody:        false,
			mockReturnRanking: nil,
		},
	}

	// Run each test
	for _, testData := range tests {
		testFramework.Run(testData.name, func(testFramework *testing.T) {
			// Setup mock
			mock := &mockDB{
				addUserSnackRankingFunc: func(ctx context.Context, user *models.UserSnackRanking) (*models.UserSnackRanking, error) {
					return testData.mockReturnRanking, testData.mockError
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
			request := httptest.NewRequest(http.MethodPost, "/snack-rankings", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, rec)

			// Call handler
			err := server.AddUserSnackRanking(ctx)
			if err != nil {
				testFramework.Errorf("AddUserSnackRanking returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectBody {
				var ranking models.UserSnackRanking
				if err := json.Unmarshal(rec.Body.Bytes(), &ranking); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if ranking.SnackID != testData.mockReturnRanking.SnackID || ranking.Rank != testData.mockReturnRanking.Rank {
					testFramework.Errorf("expected userSnackRanking %+v, got %+v", testData.mockReturnRanking, ranking)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------
// GetUserSnackRankingsByUserId
// .
// Tests
//   - success
//   - non-integer userId
//   - negative userId
//   - database error
func Test_GetUserSnackRankingsByUserId(testFramework *testing.T) {

	//---------------------------------------------
	//  CONSTANTS - for creating the userSnackRanking
	//
	TIME, _ := time.Parse(time.RFC3339, "2026-07-07T00:00:00Z")

	ingredient1 := models.Ingredient{
		ID:        1,
		Name:      "Rice Crispy Cereal",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient2 := models.Ingredient{
		ID:        2,
		Name:      "Margarine",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient3 := models.Ingredient{
		ID:        3,
		Name:      "Marshmallow",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}
	ingredient4 := models.Ingredient{
		ID:        4,
		Name:      "Vanilla",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	SNACK := models.Snack{
		ID:          1,
		Name:        "Rice Crispie Treat",
		Sweet:       true,
		Savory:      false,
		Difficulty:  2,
		RecipeUrl:   "",
		Ingredients: []models.Ingredient{ingredient1, ingredient2, ingredient3, ingredient4},
		CreatedAt:   TIME,
		UpdatedAt:   TIME,
	}

	PECAN := models.Ingredient{
		ID:        1,
		Name:      "Pecan",
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	TEAM_MULES := models.Team{
		ID:             1,
		Name:           "Mules",
		Rink:           models.RinkBairel,
		Level:          models.LevelD5,
		PrimaryColor:   "#b88907",
		SecondaryColor: "#000000",
		TernaryColor:   "#c42323",
		LogoUrl:        "",
		CreatedAt:      TIME,
		UpdatedAt:      TIME,
	}

	USER := models.User{
		ID:        1,
		FirstName: "Roger",
		LastName:  "Hogwarts",
		Email:     "r.h@gmail.com",
		Teams:     []models.Team{TEAM_MULES},
		Allergies: []models.Ingredient{PECAN},
		CreatedAt: TIME,
		UpdatedAt: TIME,
	}

	// Define the tests
	tests := []struct {
		name           string
		userId         string
		mockRankings   []models.UserSnackRanking
		expectedStatus int
		mockError      error
		expectedBody   bool // true if body should contain userSnackRankings
	}{
		{
			name:   "success",
			userId: "1",
			mockRankings: []models.UserSnackRanking{
				{SnackID: 1, Snack: SNACK, UserID: 1, User: USER, Rank: models.SnackRank10, CreatedAt: TIME, UpdatedAt: TIME},
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
			expectedBody:   true,
		},
		{
			name:           "invalid_userId",
			userId:         "bad",
			mockRankings:   nil,
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
			expectedBody:   false,
		},
		{
			name:           "negative_userId",
			userId:         "-2",
			mockRankings:   nil,
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
			expectedBody:   false,
		},
		{
			name:           "database error",
			userId:         "1",
			mockRankings:   nil,
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
				getUserSnackRankingsByUserIdFunc: func(ctx context.Context, userId int) ([]models.UserSnackRanking, error) {
					return testData.mockRankings, testData.mockError
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
			ctx.SetPath("users/:userId/snack-rankings")
			ctx.SetParamNames("userId")
			ctx.SetParamValues(testData.userId)

			// Call handler
			err := server.GetUserSnackRankingsByUserId(ctx)
			if err != nil {
				testFramework.Errorf("GetUserSnackRankingsByUserId returned error: %v", err)
			}

			// Check status
			if rec.Code != testData.expectedStatus {
				testFramework.Errorf("expected status %d, got %d", testData.expectedStatus, rec.Code)
			}

			// Check body if expected
			if testData.expectedBody {
				var rankings []models.UserSnackRanking
				if err := json.Unmarshal(rec.Body.Bytes(), &rankings); err != nil {
					testFramework.Errorf("failed to unmarshal response: %v", err)
				}
				if len(rankings) != len(testData.mockRankings) {
					testFramework.Errorf("expected %d rankings, got %d", len(testData.mockRankings), len(rankings))
				}
			}
		})
	}
}
