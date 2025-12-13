package mapper

import (
	"user-service/internal/application/common"
	"user-service/internal/domain/entities"
)

func NewUserResultFromEntityToken(user *entities.User, token string) *common.UserResult {
	if user == nil {
		return nil
	}

	return &common.UserResult{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
}
