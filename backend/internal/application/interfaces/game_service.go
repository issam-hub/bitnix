package interfaces

import (
	"bitnix-backend/internal/application/command"
	"context"
)

type GameService interface {
	CreateGame(ctx context.Context, gameCommand *command.CreateGameCommand) (*command.CreateGameCommandResult, error)
}
