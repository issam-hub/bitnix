package entities

import "github.com/google/uuid"

type Review struct {
	ReviewID uuid.UUID
	UserID   uuid.UUID
	GameID   uuid.UUID
	Rating   int
	Comment  string
}
