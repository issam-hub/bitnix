package command

import "user-service/internal/application/common"

type LoginUserCommand struct {
	Username string
	Password string
}

func NewLoginUserCommand(username, password string) *LoginUserCommand {
	return &LoginUserCommand{
		Username: username,
		Password: password,
	}
}

type LoginUserCommandResult struct {
	Result *common.LoginUserResult
}
