package response

import "github.com/google/uuid"

type RegisterUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
}
