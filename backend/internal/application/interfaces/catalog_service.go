package interfaces

import (
	"bitnix-backend/internal/application/query"
	"context"
)

type CatalogService interface {
	ListItems(ctx context.Context) (*query.ListItemsQueryResult, error)
}
