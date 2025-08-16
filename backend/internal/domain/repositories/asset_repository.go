package repositories

import (
	"bitnix-backend/internal/domain/entities"
	"context"

	"github.com/google/uuid"
)

type AssetRepository interface {
	CreateAll(ctx context.Context, assets []entities.Asset) error
	GetAllByGame(ctx context.Context, gameID uuid.UUID) ([]*entities.Asset, error)
}
