package kafkaConnection

import (
	"MessagioTestTask/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

// Config for kafka struct
type Config struct {
	Host   string `yaml:"host" env:"HOST" env-default:"broker"`
	Port   string `yaml:"port" env:"PORT" env-default:"9092"`
	Topics []struct {
		Name string `yaml:"name"`
	} `yaml:"topics"`
}

// Kafka uses to establish connection with Kafka service and send messages to queue
type Kafka struct {
	cfg     *Config
	writers map[string]*kafka.Writer
	readers map[string]*kafka.Reader
}

// New initializes Kafka struct
func New(cfg *Config) (*Kafka, error) {
	k := &Kafka{
		cfg:     cfg,
		writers: make(map[string]*kafka.Writer),
		readers: make(map[string]*kafka.Reader),
	}

	for _, topic := range cfg.Topics {
		if err := k.CreateWriter(topic.Name); err != nil {
			return nil, err
		}
	}

	return k, nil
}

// CreateWriter creates writer for topic
func (k *Kafka) CreateWriter(topic string) error {
	if _, ok := k.writers[topic]; ok {
		return fmt.Errorf("you can't create writer for topic '%s' because it has already created", topic)
	}

	k.writers[topic] = &kafka.Writer{
		Addr:                   kafka.TCP(fmt.Sprintf("%s:%s", k.cfg.Host, k.cfg.Port)),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
	return nil
}

func (k *Kafka) CloseWriters() {
	for _, writer := range k.writers {
		writer.Close()
	}
}

func (k *Kafka) CreateReader(topic string) (*kafka.Reader, error) {
	if _, ok := k.writers[topic]; ok {
		return nil, fmt.Errorf("you can't create reader for topic '%s' because it has already created", topic)
	}

	k.readers[topic] = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{fmt.Sprintf("%s:%s", k.cfg.Host, k.cfg.Port)},
		Topic:    topic,
		MaxBytes: 10e6,
	})

	return k.readers[topic], nil
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
