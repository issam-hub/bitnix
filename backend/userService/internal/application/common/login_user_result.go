package common

import (
	"github.com/google/uuid"
)

type LoginUserResult struct {
	ID    uuid.UUID
	Token string
}
