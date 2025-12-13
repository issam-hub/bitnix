package entities

import (
	"github.com/google/uuid"
)

type Cart struct {
	UserID uuid.UUID
	Items  []Game
}
