package repositories

import (
	"bitnix-backend/internal/domain/entities"
	"context"
)

type CatalogRepository interface {
	GetAll(ctx context.Context) ([]*entities.CatalogItem, error)
}
