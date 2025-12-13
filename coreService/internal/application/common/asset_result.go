package common

import (
	"bitnix-backend/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type AssetResult struct {
	ID        uuid.UUID
	Type      entities.AssetType
	URL       string
	Filename  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
