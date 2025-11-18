package mocks

import (
	"context"
	"user-service/internal/application/command"

	"github.com/stretchr/testify/mock"
)

type MockAccountsService struct {
	mock.Mock
}

func (m *MockAccountsService) Register(ctx context.Context, userCommand *command.RegisterUserCommand) (*command.RegisterUserCommandResult, error) {
	args := m.Called(ctx, userCommand)

	if result := args.Get(0); result != nil {
		return result.(*command.RegisterUserCommandResult), args.Error(1)
	}

	return nil, args.Error(1)
}
