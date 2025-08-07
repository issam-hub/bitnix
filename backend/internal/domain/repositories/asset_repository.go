package repositories

import (
	"bitnix-backend/internal/domain/entities"
	"context"
)

type AssetRepository interface {
	CreateAll(ctx context.Context, assets []entities.Asset) error
}
