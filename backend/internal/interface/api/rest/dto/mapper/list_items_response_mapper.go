package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/interface/api/rest/dto/response"
)

func ToListItemsResponse(items []*common.CatalogResult) *response.ListItemsResponse {
	var itemsResponse []response.ItemsResponse

	for _, item := range items {
		itemsResponse = append(itemsResponse, response.ItemsResponse{
			ID:        item.GameID.String(),
			Title:     item.Title,
			Price:     item.Price.AsMajorUnits(),
			Thumbnail: item.Thumbnail,
		})
	}
	return &response.ListItemsResponse{
		Items: itemsResponse,
	}
}
