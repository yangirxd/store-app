package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	log.Printf("Creating consumer for brokers: %v, topic: %s, groupID: %s", brokers, topic, groupID)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Consume(ctx context.Context, handler func([]byte) error) {
	log.Printf("Starting to consume messages from topic: %s", c.reader.Config().Topic)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Context cancelled, stopping consumer: %v", ctx.Err())
			return
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Failed to read message: %v", err)
				continue
			}
			log.Printf("Received message: %s", string(msg.Value))
			if err := handler(msg.Value); err != nil {
				log.Printf("Failed to process message: %v", err)
			}
		}
	}
}

func (c *Consumer) Close() error {
	log.Printf("Closing consumer")
	return c.reader.Close()
}
