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
