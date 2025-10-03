package services

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/application/query"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/tests/mocks"
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

func TestListItems(t *testing.T) {
	catalogRepo := new(mocks.InMemoryCatalogRepository)
	for i := 0; i < 3; i++ {
		item := entities.NewCatalogItem(
			uuid.New(),
			fmt.Sprintf("game %d", i),
			*money.NewFromFloat(30.99, "USD"),
			"thumbnail.png",
		)
		catalogRepo.Items = append(catalogRepo.Items, *item)
	}

	var expectedResult *query.ListItemsQueryResult

	for _, item := range catalogRepo.Items {
		expectedResult.Result = append(expectedResult.Result, &common.CatalogResult{
			GameID:    item.GameID,
			Title:     item.Title,
			Price:     item.Price,
			Thumbnail: item.Thumbnail,
		})
	}

	t.Run("happy case - 200", func(t *testing.T) {
		svc := NewCatalogService(catalogRepo)

		ctx := context.Background()

		queryResult, err := svc.ListItems(ctx)
		if err != nil {
			t.Errorf("error while listing items in the catalog: %v", err)
		}

		if len(queryResult.Result) != 3 {
			t.Errorf("items length doesn't match, expected %d, got %d", 3, len(queryResult.Result))
		}

		if !reflect.DeepEqual(queryResult.Result, expectedResult.Result) {
			t.Errorf("returned items doesn't match the current items in db, got %#v, want %#v", queryResult.Result, catalogRepo.Items)
		}
	})
}
