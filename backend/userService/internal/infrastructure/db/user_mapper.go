package mongodb

import (
	"time"
	"user-service/internal/domain/entities"
	"github.com/google/uuid"
)

func toDBUser(user *entities.User) *User {
	usr := User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Password:  string(user.Password.Hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
	return &usr
}

func fromDBUser(user *User) *entities.User {
	return &entities.User{
		ID:       uuid.MustParse(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}
