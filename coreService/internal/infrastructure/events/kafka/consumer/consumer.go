package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func StartConsumer(ctx context.Context, broker, topic, eventType string) <-chan Event {
	events := make(chan Event, 100)

	go func() {
		defer close(events)

		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{broker},
			GroupID:  "user-service-consumer",
			Topic:    topic,
			MinBytes: 1e3,
			MaxBytes: 10e6,
		})

		defer reader.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				m, err := reader.ReadMessage(ctx)
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					log.Printf("kafka read error: %v", err)
					time.Sleep(2 * time.Second)
					continue
				}

				var event Event

				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("invalid event: %v", err)
					continue
				}

				if event.Type == eventType {
					events <- event
				}
			}
		}
	}()

	return events
}
