package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/domain/entities"
)

func NewListItemsResultFromEntities(items []*entities.CatalogItem) []*common.CatalogResult {
	var result []*common.CatalogResult

	for _, item := range items {
		resultItem := &common.CatalogResult{
			GameID:    item.GameID,
			Title:     item.Title,
			Price:     item.Price,
			Thumbnail: item.Thumbnail,
		}
		result = append(result, resultItem)
	}
	return result
}
