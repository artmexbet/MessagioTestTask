package consumer

import (
	"MessagioTestTask/pkg/kafkaConnection"
	"MessagioTestTask/pkg/models"
	"context"
	"encoding/json"
	"fmt"
)

type IDatabase interface {
	AddMessage(models.Message) (int64, error)
}

// Config ...
type Config struct {
	ListeningTopic string `yaml:"listeningTopic" env:"LISTENING_TOPIC" env-default:"add_message"`
}

// Consumer is a struct for consuming messages from kafka
type Consumer struct {
	cfg   *Config
	queue *kafkaConnection.Kafka
	db    IDatabase
}

// New initializes Consumer struct
func New(cfg *Config, queue *kafkaConnection.Kafka, db IDatabase) *Consumer {
	return &Consumer{
		cfg:   cfg,
		queue: queue,
		db:    db,
	}
}

func (c *Consumer) Listen() error {
	reader, err := c.queue.CreateReader(c.cfg.ListeningTopic)
	if err != nil {
		return err
	}
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		var message models.Message
		if err := json.Unmarshal(m.Value, &message); err != nil {
			return err
		}

		fmt.Println(c, c.db)
		if _, err := c.db.AddMessage(message); err != nil {
			return err
		}
	}
}
