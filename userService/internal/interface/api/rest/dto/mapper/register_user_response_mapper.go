package mapper

import (
	"user-service/internal/application/common"
	"user-service/internal/interface/api/rest/dto/response"
)

func ToRegisterUserResponse(user *common.UserResult) *response.RegisterUserResponse {
	return &response.RegisterUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    user.Token,
	}
}
