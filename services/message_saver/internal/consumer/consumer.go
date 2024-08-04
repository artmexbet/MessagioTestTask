package consumer

import (
	"context"
	"github.com/nats-io/nats.go"
)

// IBroker ...
type IBroker interface {
	CreateReader(string, nats.MsgHandler) (*nats.Subscription, error)
}

type IService interface {
	HandleMessage() nats.MsgHandler
}

// Config ...
type Config struct {
}

// Consumer is a struct for consuming messages from kafka
type Consumer struct {
	cfg     *Config
	queue   IBroker
	service IService
}

// New initializes Consumer struct
func New(cfg *Config, queue IBroker, s IService) *Consumer {
	return &Consumer{
		cfg:     cfg,
		queue:   queue,
		service: s,
	}
}

func (c *Consumer) Listen(ctx context.Context, topic string) error {
	sub, err := c.queue.CreateReader(topic, c.service.HandleMessage())
	if err != nil {
		return err
	}
	defer sub.Drain()

	select {
	case <-ctx.Done():
		return nil
	}
}
