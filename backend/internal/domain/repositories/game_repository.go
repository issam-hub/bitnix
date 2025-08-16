package repositories

import (
	"bitnix-backend/internal/domain/entities"
	"context"

	"github.com/google/uuid"
)

type GameRepository interface {
	Create(ctx context.Context, game entities.Game) error
	Get(ctx context.Context, id uuid.UUID) (*entities.Game, error)
}
