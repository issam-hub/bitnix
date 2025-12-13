package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/interface/api/rest/dto/response"
)

func ToCreateGameResponse(game *common.GameResult) *response.CreateGameResponse {
	return &response.CreateGameResponse{
		ID: game.ID,
	}
}
