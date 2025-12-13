package common

import (
	"github.com/google/uuid"
)

type UserResult struct {
	ID       uuid.UUID
	Username string
	Email    string
	Token    string
}
