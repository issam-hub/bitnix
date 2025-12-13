package services

import (
	"context"
	"fmt"
	"testing"
	"user-service/internal/application/command"
	"user-service/internal/tests/mocks"
)

func TestRegister(t *testing.T) {
	usersRepo := new(mocks.InMemoryUsersRepository)
	failingUsersRepo := new(mocks.FailingInMemoryUsersRepository)

	cmd := command.NewRegisterUserCommand(
		"test_test",
		"test@example.com",
		"password123",
		"developer",
	)
	t.Run("happy case", func(t *testing.T) {
		svc := NewAccountsService(usersRepo)

		ctx := context.Background()

		result, err := svc.Register(ctx, cmd)
		if err != nil {
			t.Errorf("error while registering user: %#v", err)
		}

		if len(usersRepo.Users) == 0 {
			t.Errorf("error while registering user: want %d length, got %d length", 1, len(usersRepo.Users))
		}

		gotUsername, gotPassword, _ := DecodeBasicToken(result.Result.Token)

		if gotUsername != "test_test" || gotPassword != "password123" {
			t.Errorf("error while registering user: token mismatch")
		}
	})

	t.Run("sad case", func(t *testing.T) {
		svc := NewAccountsService(failingUsersRepo)

		ctx := context.Background()

		if _, err := svc.Register(ctx, cmd); err == nil || err.Error() != "error while creating user, coming from mongodb" {
			t.Errorf("expected nil or 'error while creating user, coming from mongodb' error, got: %v", err)
		}

		if len(failingUsersRepo.Users) != 0 {
			t.Errorf("error while registering user: want %d length, got %d length", 0, len(failingUsersRepo.Users))
		}
	})
}

func TestLogin(t *testing.T) {
	usersRepo := new(mocks.InMemoryUsersRepository)

	cmd := command.NewLoginUserCommand(
		"test_test",
		"password123",
	)

	createCmd := command.NewRegisterUserCommand(
		"test_test",
		"test@example.com",
		"password123",
		"developer",
	)
	t.Run("happy case", func(t *testing.T) {
		svc := NewAccountsService(usersRepo)

		ctx := context.Background()

		if _, err := svc.Register(ctx, createCmd); err != nil {
			t.Errorf("error while creating user: %#v", err)
		}

		result, err := svc.Login(ctx, cmd)
		if err != nil {
			t.Errorf("error while logging in user: %#v", err)
		}

		expectedToken := "dGVzdF90ZXN0OnBhc3N3b3JkMTIz"

		if result.Result.Token != fmt.Sprintf("Basic %s", expectedToken) {
			t.Errorf("error while logging in user: invalid credentials")
		}
	})
}
