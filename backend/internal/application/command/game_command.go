package command

import (
	"bitnix-backend/internal/application/common"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type CreateGameCommand struct {
	Title       string
	Description string
	Price       money.Money
	DeveloperID uuid.UUID
	ReleaseDate time.Time
	Assets      []uuid.UUID
}

func NewCreateGameCommand(title string, description string, price money.Money, developerID uuid.UUID, releaseDate time.Time, assets []uuid.UUID) *CreateGameCommand {
	return &CreateGameCommand{
		Title:       title,
		Description: description,
		Price:       price,
		DeveloperID: developerID,
		ReleaseDate: releaseDate,
		Assets:      assets,
	}
}

type CreateGameCommandResult struct {
	Result *common.GameResult
}
