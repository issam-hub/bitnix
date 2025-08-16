package resttest

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/application/query"
	"bitnix-backend/internal/domain/apperrors"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/interface/api/rest"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ?info: The mock essentially lets you say: "Pretend the service did its job successfully/unsuccessfully, and let me test how my controller handles that result."

func TestCreateGame(t *testing.T) {
	e := echo.New()

	t.Run("happy case - 201", func(t *testing.T) {
		mockSvc := new(MockGameService)

		devID := uuid.New()
		assets := []map[string]string{
			{
				"type":     "download",
				"url":      "https://downloadMe.com",
				"filename": "downloadFile",
			},
		}
		reqBody := map[string]any{
			"title":        "hollow knight",
			"description":  "hollow knight game",
			"price":        19.99,
			"release_date": "2025-10-06",
			"developer_id": devID.String(),
			"assets":       assets,
		}

		reqBodyBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/game", bytes.NewReader(reqBodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewgameController(e, mockSvc)
		releaseDate, _ := time.Parse("2006-01-02", "2025-10-06")
		gameID := uuid.New()
		createdAt := time.Now()
		assetID := uuid.New()
		createGameCommandResult := &command.CreateGameCommandResult{
			Result: &common.GameResult{
				ID:          gameID,
				Title:       "hollow knight",
				Description: "hollow knight game",
				Price:       *money.NewFromFloat(19.99, "USD"),
				DeveloperID: devID,
				ReleaseDate: releaseDate,
				Assets: []entities.Asset{
					{
						ID:       assetID,
						Type:     entities.DownloadFile,
						GameID:   gameID,
						URL:      "https://downloadMe.com",
						Filename: "downloadFile",
					},
				},
				CreatedAt: createdAt,
			},
		}

		mockSvc.On("CreateGame", mock.Anything, mock.AnythingOfType("*command.CreateGameCommand")).Return(createGameCommandResult, nil)

		err := ctrl.CreateGameController(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rec.Code)

		expectedResponseBody := map[string]any{
			"id": gameID.String(),
		}

		var actualResponseBody map[string]any

		err = json.Unmarshal(rec.Body.Bytes(), &actualResponseBody)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		assert.Equal(t, expectedResponseBody, actualResponseBody)

		mockSvc.AssertExpectations(t)
	})

	t.Run("sad case - 400", func(t *testing.T) {
		mockSvc := new(MockGameService)

		invalidReqBody := map[string]any{
			"title":        "",
			"description":  "",
			"price":        0,
			"release_date": "2025-10-06",
			"developer_id": "",
			"assets":       []map[string]string{},
		}

		reqBodyBytes, _ := json.Marshal(invalidReqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/game", bytes.NewReader(reqBodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewgameController(e, mockSvc)

		err := ctrl.CreateGameController(c)

		assert.Error(t, err)

		httpErr, _ := err.(*echo.HTTPError)

		assert.Equal(t, http.StatusBadRequest, httpErr.Code)

		expectedErrors := map[string]string{
			"title":        "title is required and cannot be empty",
			"description":  "description is required and cannot be empty",
			"price":        "price cannot be zero or negative",
			"developer_id": "developer ID is required",
			"assets":       "it should be at least one asset",
		}

		assert.Equal(t, expectedErrors, httpErr.Message)

	})

	t.Run("sad case - 500", func(t *testing.T) {
		mockSvc := new(MockGameService)

		devID := uuid.New()
		assets := []map[string]string{
			{
				"type":     "download",
				"url":      "https://downloadMe.com",
				"filename": "downloadFile",
			},
		}
		reqBody := map[string]any{
			"title":        "hollow knight",
			"description":  "hollow knight game",
			"price":        19.99,
			"release_date": "2025-10-06",
			"developer_id": devID.String(),
			"assets":       assets,
		}

		reqBodyBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/game", bytes.NewReader(reqBodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewgameController(e, mockSvc)

		mockSvc.On("CreateGame", mock.Anything, mock.AnythingOfType("*command.CreateGameCommand")).Return(nil, errors.New("Internal server error trigger"))

		err := ctrl.CreateGameController(c)

		assert.Error(t, err)

		httpErr, _ := err.(*echo.HTTPError)
		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)

		expectedError := "Internal Server Error"

		assert.Equal(t, expectedError, httpErr.Message)

		mockSvc.AssertExpectations(t)
	})
}

func TestGetGame(t *testing.T) {
	e := echo.New()

	t.Run("happy case - 200", func(t *testing.T) {
		mockSvc := new(MockGameService)

		gameID := uuid.New()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/game/%s", gameID.String()), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(gameID.String())

		ctrl := rest.NewgameController(e, mockSvc)

		releaseDate, _ := time.Parse("2006-01-02", "2025-10-06")
		createdAt := time.Now()
		assetID := uuid.New()
		devID := uuid.New()
		getGameQueryResult := &query.GameQueryResult{
			Result: &common.GameResult{
				ID:          gameID,
				Title:       "hollow knight",
				Description: "hollow knight game",
				Price:       *money.NewFromFloat(19.99, "USD"),
				DeveloperID: devID,
				ReleaseDate: releaseDate,
				Assets: []entities.Asset{
					{
						ID:       assetID,
						Type:     entities.DownloadFile,
						GameID:   gameID,
						URL:      "https://downloadMe.com",
						Filename: "downloadFile",
					},
				},
				CreatedAt: createdAt,
			},
		}

		mockSvc.On("GetGame", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(getGameQueryResult, nil)

		err := ctrl.GetGameController(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		var expectedAssets []any

		for _, asset := range getGameQueryResult.Result.Assets {
			expectedAssets = append(expectedAssets, map[string]any{
				"id":       asset.ID.String(),
				"type":     string(asset.Type),
				"filename": asset.Filename,
				"url":      asset.URL,
			})
		}

		expectedResponseBody := map[string]any{
			"id":           gameID.String(),
			"title":        getGameQueryResult.Result.Title,
			"description":  getGameQueryResult.Result.Description,
			"price":        getGameQueryResult.Result.Price.AsMajorUnits(),
			"developer_id": getGameQueryResult.Result.DeveloperID.String(),
			"release_date": getGameQueryResult.Result.ReleaseDate.String(),
			"assets":       expectedAssets,
		}

		var actualResponseBody map[string]any

		err = json.Unmarshal(rec.Body.Bytes(), &actualResponseBody)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		assert.Equal(t, expectedResponseBody, actualResponseBody)

		mockSvc.AssertExpectations(t)
	})

	t.Run("sade case - 400", func(t *testing.T) {
		mockSvc := new(MockGameService)

		gameID := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/game/%s", gameID), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(gameID)

		ctrl := rest.NewgameController(e, mockSvc)

		err := ctrl.GetGameController(c)

		assert.Error(t, err)

		httpErr, _ := err.(*echo.HTTPError)
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)

		expectedError := "invalid game ID format"

		assert.Equal(t, expectedError, httpErr.Message)
	})

	t.Run("sad case - 404", func(t *testing.T) {
		mockSvc := new(MockGameService)

		gameID := uuid.New()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/game/%s", gameID.String()), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(gameID.String())

		ctrl := rest.NewgameController(e, mockSvc)

		mockSvc.On("GetGame", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, apperrors.ErrGameNotFound)

		err := ctrl.GetGameController(c)

		assert.Error(t, err)

		httpErr, _ := err.(*echo.HTTPError)
		assert.Equal(t, http.StatusNotFound, httpErr.Code)

		assert.Equal(t, apperrors.ErrGameNotFound.Error(), httpErr.Message)

		mockSvc.AssertExpectations(t)
	})
}
