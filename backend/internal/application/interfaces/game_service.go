package interfaces

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/query"
	"context"

	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(ctx context.Context, gameCommand *command.CreateGameCommand) (*command.CreateGameCommandResult, error)
	GetGame(ctx context.Context, id uuid.UUID) (*query.GameQueryResult, error)
}
