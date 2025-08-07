package resttest

import (
	"bitnix-backend/internal/application/command"
	"context"

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
