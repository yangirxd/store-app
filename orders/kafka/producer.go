package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Produce(ctx context.Context, topic string, message []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: message,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
