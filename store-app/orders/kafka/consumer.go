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
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
	}
}

func (c *Consumer) Consume(ctx context.Context, handler func([]byte) error) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			continue
		}

		if err := handler(msg.Value); err != nil {
			log.Printf("Failed to process message: %v", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
