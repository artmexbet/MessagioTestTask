package main

import (
	"MessagioTestTask/pkg/logger/sl"
	"MessagioTestTask/pkg/natsConn"
	"MessagioTestTask/services/gateway/internal/router"
	"MessagioTestTask/services/gateway/internal/service"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

// Config is a global service config structure
type Config struct {
	RouterConfig router.Config `yaml:"router" env-prefix:"ROUTER_"`
	//KafkaConfig   kafkaConnection.Config `yaml:"kafka" env-prefix:"KAFKA_"`
	NatsConfig    natsConn.Config `yaml:"nats" env-prefix:"NATS_"`
	ServiceConfig service.Config  `yaml:"service" env-prefix:"SERVICE_"`
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

	//k, err := kafkaConnection.New(&cfg.KafkaConfig)
	//if err != nil {
	//	log.Fatalln(err) // Cannot connect to Kafka
	//}
	//defer k.CloseWriters()

	n, err := natsConn.New(&cfg.NatsConfig)
	if err != nil {
		log.Fatalln(err)
	}

	svc, err := service.New(&cfg.ServiceConfig, n)
	if err != nil {
		log.Fatalln(err)
	}

	r := router.New(&cfg.RouterConfig, svc)
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
