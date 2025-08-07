package repositories

import (
	"bitnix-backend/internal/domain/entities"
	"context"
)

type GameRepository interface {
	Create(ctx context.Context, game entities.Game) error
}
