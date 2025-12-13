package publisher

import (
	"time"

	"github.com/google/uuid"
)

type OutboxEvent struct {
	ID          uuid.UUID  `bson:"id"`
	AggregateID uuid.UUID  `bson:"aggregate_id"`
	EventType   string     `bson:"event_type"`
	Payload     string     `bson:"payload"`
	Published   bool       `bson:"published"`
	CreatedAt   time.Time  `bson:"created_at"`
	PublishedAt *time.Time `bson:"published_at"`
}
