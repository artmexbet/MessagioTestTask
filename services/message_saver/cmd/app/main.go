package main

import (
	"MessagioTestTask/pkg/db"
	"MessagioTestTask/pkg/kafkaConnection"
	"MessagioTestTask/pkg/logger/sl"
	"MessagioTestTask/services/message_saver/internal/consumer"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"log/slog"
)

// Config ...
type Config struct {
	KafkaConfig    kafkaConnection.Config `yaml:"kafka" env-prefix:"KAFKA_"`
	ConsumerConfig consumer.Config        `yaml:"consumer" env-prefix:"CONSUMER_"`
	DBConfig       db.Config              `yaml:"db" env-prefix:"DB_"`
}

// readConfig gets config from file filename
func readConfig(filename string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	sl.SetupLogger("local")

	cfg, err := readConfig("./config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	k, err := kafkaConnection.New(&cfg.KafkaConfig)
	if err != nil {
		log.Fatalln(err) // Cannot connect to Kafka
	}

	DB, err := db.New(&cfg.DBConfig)
	if err != nil {
		slog.With("module", "message_saver.consumer").With("raizer", "db").Error(err.Error())
	}

	c := consumer.New(&cfg.ConsumerConfig, k, DB)
	err = c.Listen()
	if err != nil {
		slog.With("module", "message_saver.consumer").Error(err.Error())
	}
}
