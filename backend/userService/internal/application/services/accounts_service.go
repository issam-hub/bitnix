package services

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"
	"user-service/internal/application/command"
	"user-service/internal/application/mapper"
	"user-service/internal/domain/entities"
	"user-service/internal/domain/repositories"
)

func EncodeBasicToken(username, password string) string {
	raw := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(raw))
}

func DecodeBasicToken(token string) (string, string, error) {
	if after, ok := strings.CutPrefix(token, "Basic "); ok {
		token = after
	}
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(token))
	if err != nil {
		return "", "", err
	}
	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return "", "", errors.New("invalid basic auth payload")
	}
	return parts[0], parts[1], nil
}

type AccountsService struct {
	usersRepo repositories.UsersRepository
}

func NewAccountsService(usersRepo repositories.UsersRepository) *AccountsService {
	return &AccountsService{
		usersRepo: usersRepo,
	}
}

func (s AccountsService) Register(ctx context.Context, userCommand *command.RegisterUserCommand) (*command.RegisterUserCommandResult, error) {
	user := entities.NewUser(
		userCommand.Username,
		userCommand.Email,
		userCommand.Password,
		userCommand.Role,
	)

	if err := s.usersRepo.Create(ctx, *user); err != nil {
		return nil, err
	}

	token := EncodeBasicToken(userCommand.Username, userCommand.Password)

	commandResult := mapper.NewUserResultFromEntityToken(user, token)

	return &command.RegisterUserCommandResult{
		Result: commandResult,
	}, nil

}
