package main

import (
	"MessagioTestTask/pkg/kafkaConnector"
	"MessagioTestTask/pkg/logger/sl"
	"MessagioTestTask/services/gateway/internal/router"
	"MessagioTestTask/services/gateway/internal/service"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

// Config is a global project config structure
type Config struct {
	RouterConfig  router.Config         `yaml:"router" env-prefix:"ROUTER_"`
	KafkaConfig   kafkaConnector.Config `yaml:"kafka" env-prefix:"KAFKA_"`
	ServiceConfig service.Config        `yaml:"service" env-prefix:"SERVICE_"`
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

	k, err := kafkaConnector.New(&cfg.KafkaConfig)
	if err != nil {
		log.Fatalln(err) // Cannot connect to Kafka
	}

	svc, err := service.New(&cfg.ServiceConfig, k)
	if err != nil {
		log.Fatalln(err)
	}

	r := router.New(&cfg.RouterConfig, svc)
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
