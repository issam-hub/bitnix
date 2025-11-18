package repositories

import (
	"context"
	"user-service/internal/domain/entities"
)

type UsersRepository interface {
	Create(ctx context.Context, user entities.User) error
}
