package services

import (
	"bitnix-backend/internal/application/mapper"
	"bitnix-backend/internal/application/query"
	"bitnix-backend/internal/domain/repositories"
	"context"
)

type CatalogService struct {
	catalogRepository repositories.CatalogRepository
}

func NewCatalogService(catalogRepo repositories.CatalogRepository) *CatalogService {
	return &CatalogService{
		catalogRepository: catalogRepo,
	}
}

func (s CatalogService) ListItems(ctx context.Context) (*query.ListItemsQueryResult, error) {
	items, err := s.catalogRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	queryResult := mapper.NewListItemsResultFromEntities(items)
	return &query.ListItemsQueryResult{
		Result: queryResult,
	}, nil
}
