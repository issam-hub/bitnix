package entities

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Game struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       money.Money
	Developer   User
	ReleaseDate time.Time
	Assets      []Asset
}

func NewGame(title string, description string, price money.Money, developer User, releaseDate time.Time, assets []Asset) *Game {
	return &Game{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Price:       price,
		Developer:   developer,
		ReleaseDate: releaseDate,
		Assets:      assets,
	}
}
