package common

import (
	"bitnix-backend/internal/domain/entities"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type GameResult struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       money.Money
	DeveloperID uuid.UUID
	ReleaseDate time.Time
	Assets      []entities.Asset
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
