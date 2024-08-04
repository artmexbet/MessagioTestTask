package main

import (
	"MessagioTestTask/pkg/db"
	"MessagioTestTask/pkg/logger/sl"
	"MessagioTestTask/pkg/natsConn"
	"MessagioTestTask/services/message_saver/internal/consumer"
	"MessagioTestTask/services/message_saver/internal/service"
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"log/slog"
)

// Config ...
type Config struct {
	NatsConfig     natsConn.Config `yaml:"nats" env-prefix:"NATS_"`
	ConsumerConfig consumer.Config `yaml:"consumer" env-prefix:"CONSUMER_"`
	ServiceConfig  service.Config  `yaml:"service" env-prefix:"SERVICE_"`
	DBConfig       db.Config       `yaml:"db" env-prefix:"DB_"`
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

	n, err := natsConn.New(&cfg.NatsConfig)
	if err != nil {
		slog.With("module", "message_saver.nats").Error(err.Error())
	}

	DB, err := db.New(&cfg.DBConfig)
	if err != nil {
		slog.With("module", "message_saver.consumer").With("raizer", "db").Error(err.Error())
	}

	s, err := service.New(&cfg.ServiceConfig, DB, n)

	c := consumer.New(&cfg.ConsumerConfig, n, s)
	ctx := context.Background()
	err = c.Listen(ctx, "add_message")
	if err != nil {
		slog.With("module", "message_saver.consumer").Error(err.Error())
	}
}
