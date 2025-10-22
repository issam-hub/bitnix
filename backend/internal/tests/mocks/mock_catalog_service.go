package mocks

import (
	"bitnix-backend/internal/application/query"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCatalogService struct {
	mock.Mock
}

func (m *MockCatalogService) ListItems(ctx context.Context) (*query.ListItemsQueryResult, error) {
	args := m.Called(ctx)

	if result := args.Get(0); result != nil {
		return result.(*query.ListItemsQueryResult), args.Error(1)
	}

	return nil, args.Error(1)
}
