package postgres

import (
	"bitnix-backend/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       float64
	DeveloperID uuid.UUID
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Version     int32
}

type Asset struct {
	ID        uuid.UUID
	GameID    uuid.UUID
	Type      entities.AssetType
	URL       string
	Filename  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int32
}
