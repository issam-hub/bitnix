package postgres

import (
	"bitnix-backend/internal/domain/entities"

	"github.com/Rhymond/go-money"
)

func toDBGame(game *entities.Game) *Game {
	return &Game{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
		Price:       game.Price.AsMajorUnits(),
		DeveloperID: game.Developer.ID,
		ReleaseDate: game.ReleaseDate,
	}
}

func fromDBGame(game *Game) *entities.Game {
	return &entities.Game{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
		Price:       *money.NewFromFloat(game.Price, "USD"),
		ReleaseDate: game.ReleaseDate,
	}
}
