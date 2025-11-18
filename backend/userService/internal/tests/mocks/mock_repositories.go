package mocks

import (
	"context"
	"errors"
	"user-service/internal/domain/entities"
)

type InMemoryUsersRepository struct {
	Users []entities.User
}

func (mu *InMemoryUsersRepository) Create(ctx context.Context, user entities.User) error {
	mu.Users = append(mu.Users, user)
	return nil
}

type FailingInMemoryUsersRepository struct {
	Users []entities.User
}

func (fmu *FailingInMemoryUsersRepository) Create(ctx context.Context, user entities.User) error {
	return errors.New("error while creating user, coming from mongodb")
}
