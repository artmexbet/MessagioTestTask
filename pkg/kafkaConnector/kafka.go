package kafkaConnector

import (
	"MessagioTestTask/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

// IQueue interface for kafka
type IQueue interface {
	SendEvent(models.Message, string) error
}

// Config for kafka struct
type Config struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"PORT" env-default:"9092"`
}

// Kafka uses to establish connection with Kafka service and send messages to queue
type Kafka struct {
	cfg     *Config
	writers map[string]*kafka.Writer
}

// New initializes Kafka struct
func New(cfg *Config) (*Kafka, error) {
	k := &Kafka{
		cfg:     cfg,
		writers: make(map[string]*kafka.Writer),
	}

	return k, nil
}

// CreateWriter creates writer for topic
func (k *Kafka) CreateWriter(topic string) error {
	k.writers[topic] = &kafka.Writer{
		Addr:                   kafka.TCP(fmt.Sprintf("%s:%s", k.cfg.Host, k.cfg.Port)),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
	return nil
}

// SendEvent sends event to queue with some info
func (k *Kafka) SendEvent(message models.Message, topic string) error {
	if _, ok := k.writers[topic]; !ok {
		if err := k.CreateWriter(topic); err != nil {
			return err
		}
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	writer := k.writers[topic]

	err = writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(message.Id),
			Value: messageBytes,
		},
	)

	return err
}
