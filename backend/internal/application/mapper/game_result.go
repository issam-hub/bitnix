package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/domain/entities"
)

func NewGameResultFromEntity(game *entities.Game) *common.GameResult {
	if game == nil {
		return nil
	}
	return &common.GameResult{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
		Price:       game.Price,
		DeveloperID: game.Developer.ID,
		ReleaseDate: game.ReleaseDate,
		Assets:      game.Assets,
	}
}
