package mocks

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/query"
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGameService struct {
	mock.Mock
}

func (m *MockGameService) CreateGame(ctx context.Context, gameCommand *command.CreateGameCommand) (*command.CreateGameCommandResult, error) {
	args := m.Called(ctx, gameCommand)

	if result := args.Get(0); result != nil {
		return result.(*command.CreateGameCommandResult), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockGameService) GetGame(ctx context.Context, id uuid.UUID) (*query.GameQueryResult, error) {
	args := m.Called(ctx, id)

	if result := args.Get(0); result != nil {
		return result.(*query.GameQueryResult), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockGameService) UploadAssets(ctx context.Context, assetsCommand *command.UploadAssetsCommand) (*command.UploadAssetsCommandResult, error) {
	args := m.Called(ctx, assetsCommand)

	if result := args.Get(0); result != nil {
		return result.(*command.UploadAssetsCommandResult), args.Error(1)
	}

	return nil, args.Error(1)
}
