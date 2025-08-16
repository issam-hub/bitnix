package command

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/domain/entities"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type AssetDetail struct {
	Type     entities.AssetType
	URL      string
	Filename string
}

type CreateGameCommand struct {
	Title       string
	Description string
	Price       money.Money
	DeveloperID uuid.UUID
	ReleaseDate time.Time
	Assets      []AssetDetail
}

func NewCreateGameCommand(title string, description string, price money.Money, developerID uuid.UUID, releaseDate time.Time, assets []AssetDetail) *CreateGameCommand {
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
