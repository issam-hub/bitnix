package common

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type CatalogResult struct {
	GameID    uuid.UUID
	Title     string
	Price     money.Money
	Thumbnail string
}
