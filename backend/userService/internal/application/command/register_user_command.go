package command

import "user-service/internal/application/common"

type RegisterUserCommand struct {
	Username string
	Email    string
	Password string
	Role     string
}

func NewRegisterUserCommand(username, email, password, role string) *RegisterUserCommand {
	return &RegisterUserCommand{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

type RegisterUserCommandResult struct {
	Result *common.UserResult
}
