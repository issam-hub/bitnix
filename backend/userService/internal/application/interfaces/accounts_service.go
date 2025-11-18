package interfaces

import (
	"context"
	"user-service/internal/application/command"
)

type AccountsService interface {
	Register(ctx context.Context, userCommand *command.RegisterUserCommand) (*command.RegisterUserCommandResult, error)
}
