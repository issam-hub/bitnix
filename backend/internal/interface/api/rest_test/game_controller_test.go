package resttest

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/interface/api/rest"
	"bytes"
	"encoding/json"
	"errors"
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

	mockSvc := new(MockGameService)

	t.Run("happy case - 201", func(t *testing.T) {
		devID := uuid.New()
		assetsIDs := []string{
			uuid.New().String(),
		}
		reqBody := map[string]any{
			"title":        "hollow knight",
			"description":  "hollow knight game",
			"price":        19.99,
			"release_date": "2025-10-06",
			"developer_id": devID.String(),
			"assets":       assetsIDs,
		}

		reqBodyBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/game", bytes.NewReader(reqBodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewgameController(e, mockSvc)
		releaseDate, _ := time.Parse("2006-01-02", "2025-10-06")
		var parsedAssetsIDs []uuid.UUID
		for _, assetID := range assetsIDs {
			parsedAssetID, _ := uuid.Parse(assetID)
			parsedAssetsIDs = append(parsedAssetsIDs, parsedAssetID)
		}
		gameID := uuid.New()
		createdAt := time.Now()
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
						ID:       parsedAssetsIDs[0],
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
		invalidReqBody := map[string]any{
			"title":        "",
			"description":  "",
			"price":        0,
			"release_date": "2025-10-06",
			"developer_id": "",
			"assets":       []string{},
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
		devID := uuid.New()
		assetsIDs := []string{
			uuid.New().String(),
		}
		reqBody := map[string]any{
			"title":        "hollow knight",
			"description":  "hollow knight game",
			"price":        19.99,
			"release_date": "2025-10-06",
			"developer_id": devID.String(),
			"assets":       assetsIDs,
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
