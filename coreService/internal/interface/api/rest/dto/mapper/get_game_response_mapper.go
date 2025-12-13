package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/interface/api/rest/dto/response"
)

func ToGetGameResponse(game *common.GameResult) *response.GetGameResponse {
	var responseAssets []response.AssetResponse
	for _, asset := range game.Assets {
		responseAssets = append(responseAssets, response.AssetResponse{
			ID:       asset.ID,
			Type:     string(asset.Type),
			Filename: asset.Filename,
			URL:      asset.URL,
		})
	}
	return &response.GetGameResponse{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
		Price:       game.Price.AsMajorUnits(),
		DeveloperID: game.DeveloperID.String(),
		ReleaseDate: game.ReleaseDate.String(),
		Assets:      responseAssets,
	}
}
