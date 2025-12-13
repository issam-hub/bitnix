package publisher

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func StartPoller(db *mongo.Database, broker string, topic string, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			publishPending(db, broker, topic)
		}
	}()
}

func publishPending(db *mongo.Database, broker, topic string) {
	var events []OutboxEvent
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := db.Collection("outbox-table")

	cursor, err := coll.Find(ctx, bson.D{{"published", false}})
	if err != nil {
		log.Printf("outbox query error: %v", err)
		return
	}

	if err = cursor.All(context.TODO(), &events); err != nil {
		log.Printf("outbox query error: %v", err)
		return
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

		_, err = coll.UpdateOne(ctx, bson.D{{"id", event.ID}}, bson.D{{"$set", bson.D{{"published", true}, {"published_at", time.Now()}}}})
		if err != nil {
			log.Printf("failed to mark outbox published %d: %v", event.ID, err)
			continue
		}
	}

}
