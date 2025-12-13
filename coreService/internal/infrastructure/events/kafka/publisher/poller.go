package publisher

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func StartPoller(db *sql.DB, broker string, topic string, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			publishPending(db, broker, topic)
		}
	}()
}

func publishPending(db *sql.DB, broker, topic string) {
	var events []OutboxEvent
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `SELECT id, aggregate_id, event_type, payload, published, created_at, published_at FROM outbox-table WHERE published = false`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("outbox query error: %v", err)
		return
	}

	for rows.Next() {
		var event OutboxEvent
		err := rows.Scan(
			&event.ID,
			&event.AggregateID,
			&event.EventType,
			&event.Payload,
			&event.Published,
			&event.CreatedAt,
			&event.PublishedAt,
		)

		if err != nil {
			log.Printf("outbox query error: %v", err)
			return
		}

		events = append(events, event)
	}

	if len(events) == 0 {
		return
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for _, event := range events {
		msg := kafka.Message{
			Key:     []byte(event.AggregateID.String()),
			Value:   []byte(event.Payload),
			Headers: []kafka.Header{{Key: "eventType", Value: []byte(event.EventType)}},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err := writer.WriteMessages(ctx, msg)
		cancel()

		if err != nil {
			log.Printf("failed to write event %d: %v", event.ID, err)
			continue
		}

		query = `UPDATE outbox-table SET published = true and published_at = $1 WHERE id = $2`

		args := []any{
			time.Now(),
			event.ID,
		}

		result, err := db.ExecContext(ctx, query, args...)
		if err != nil {
			log.Printf("failed to mark outbox published %d: %v", event.ID, err)
			continue
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("failed to mark outbox published %d: %v", event.ID, err)
			continue
		}

		if rowsAffected == 0 {
			log.Printf("failed to mark outbox published %d: %v", event.ID, err)
			continue
		}
	}

}
