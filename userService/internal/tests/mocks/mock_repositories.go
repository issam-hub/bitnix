package mocks

import (
	"context"
	"errors"
	"user-service/internal/domain/apperrors"
	"user-service/internal/domain/entities"
)

type InMemoryUsersRepository struct {
	Users []entities.User
}

func (mu *InMemoryUsersRepository) Create(ctx context.Context, user entities.User) error {
	mu.Users = append(mu.Users, user)
	return nil
}

func (mu *InMemoryUsersRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	for _, u := range mu.Users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, apperrors.ErrUserNotFound
}

type FailingInMemoryUsersRepository struct {
	Users []entities.User
}

func (fmu *FailingInMemoryUsersRepository) Create(ctx context.Context, user entities.User) error {
	return errors.New("error while creating user, coming from mongodb")
}

func (mu *FailingInMemoryUsersRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	return nil, nil
}
