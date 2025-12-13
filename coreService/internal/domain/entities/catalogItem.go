package entities

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type CatalogItem struct {
	GameID    uuid.UUID
	Title     string
	Price     money.Money
	Thumbnail string
}

func NewCatalogItem(gameID uuid.UUID, title string, price money.Money, thumbnail string) *CatalogItem {
	return &CatalogItem{
		GameID:    gameID,
		Title:     title,
		Price:     price,
		Thumbnail: thumbnail,
	}
}
