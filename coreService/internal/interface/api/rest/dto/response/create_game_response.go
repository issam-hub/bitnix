package response

import (
	"github.com/google/uuid"
)

type CreateGameResponse struct {
	ID uuid.UUID `json:"id"`
}
