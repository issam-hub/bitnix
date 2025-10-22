package resttest

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/application/query"
	"bitnix-backend/internal/interface/api/rest"
	"bitnix-backend/internal/tests/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListCatalog(t *testing.T) {
	e := echo.New()

	t.Run("happy case - 200", func(t *testing.T) {
		mockSvc := new(mocks.MockCatalogService)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/catalog", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewCatalogController(e, mockSvc)

		itemsIDs := []uuid.UUID{
			uuid.New(),
			uuid.New(),
		}

		listItemsQueryResult := &query.ListItemsQueryResult{
			Result: []*common.CatalogResult{
				{
					GameID:    itemsIDs[0],
					Title:     "hollow knight",
					Price:     *money.NewFromFloat(10.99, "USD"),
					Thumbnail: "thumbnail.png",
				},
				{
					GameID:    itemsIDs[1],
					Title:     "hollow knight: silksong",
					Price:     *money.NewFromFloat(19.99, "USD"),
					Thumbnail: "thumbnail.png",
				},
			},
		}

		mockSvc.On("ListItems", mock.Anything).Return(listItemsQueryResult, nil)

		err := ctrl.ListItemsController(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		expectedResponseBody := map[string]any{
			"items": []any{
				map[string]any{
					"id":        itemsIDs[0].String(),
					"title":     listItemsQueryResult.Result[0].Title,
					"price":     listItemsQueryResult.Result[0].Price.AsMajorUnits(),
					"thumbnail": listItemsQueryResult.Result[0].Thumbnail,
				},
				map[string]any{
					"id":        itemsIDs[1].String(),
					"title":     listItemsQueryResult.Result[1].Title,
					"price":     listItemsQueryResult.Result[1].Price.AsMajorUnits(),
					"thumbnail": listItemsQueryResult.Result[1].Thumbnail,
				},
			},
		}

		var actualResponseBody map[string]any

		err = json.Unmarshal(rec.Body.Bytes(), &actualResponseBody)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		assert.Equal(t, expectedResponseBody, actualResponseBody)

		mockSvc.AssertExpectations(t)
	})

	t.Run("sad case - 500", func(t *testing.T) {
		mockSvc := new(mocks.MockCatalogService)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/catalog", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewCatalogController(e, mockSvc)

		mockSvc.On("ListItems", mock.Anything).Return(nil, errors.New("internal server error"))

		err := ctrl.ListItemsController(c)

		assert.Error(t, err)

		httpErr, _ := err.(*echo.HTTPError)

		assert.Equal(t, http.StatusInternalServerError, httpErr.Code)

		expectedError := "Internal Server Error"

		assert.Equal(t, expectedError, httpErr.Message)

		mockSvc.AssertExpectations(t)
	})
}
