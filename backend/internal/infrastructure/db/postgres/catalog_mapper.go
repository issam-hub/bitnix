package postgres

import (
	"bitnix-backend/internal/domain/entities"

	"github.com/Rhymond/go-money"
)

func fromDbCatalogItem(item *CatalogItem) *entities.CatalogItem {
	return &entities.CatalogItem{
		GameID:    item.GameID,
		Title:     item.Title,
		Price:     *money.NewFromFloat(item.Price, "USD"),
		Thumbnail: item.Thumbnail,
	}
}
