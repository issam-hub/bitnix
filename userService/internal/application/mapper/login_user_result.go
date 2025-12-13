package mapper

import (
	"user-service/internal/application/common"
	"user-service/internal/domain/entities"
)

func NewLoginUserResultFromEntity(user *entities.User, token string) *common.LoginUserResult {
	if user == nil {
		return nil
	}

	return &common.LoginUserResult{
		ID:    user.ID,
		Token: token,
	}
}
