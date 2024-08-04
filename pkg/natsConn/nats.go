package natsConn

import (
	"MessagioTestTask/pkg/models"
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type Config struct {
	Host string `yaml:"host" env-prefix:"HOST" env-default:"localhost"`
	Port string `yaml:"port" env-prefix:"PORT" env-default:"4222"`
}

type Nats struct {
	cfg  *Config
	conn *nats.Conn
}

func New(cfg *Config) (*Nats, error) {
	n := &Nats{
		cfg: cfg,
	}

	conn, err := nats.Connect(cfg.Host + ":" + cfg.Port)
	if err != nil {
		return nil, err
	}
	n.conn = conn

	return n, nil
}

func (n *Nats) CreateReader(topic string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return n.conn.Subscribe(topic, handler)
}

func (n *Nats) SendMessage(message models.Message, topic string) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return n.conn.Publish(topic, bytes)
}

func (n *Nats) SendMsgInfo(msgInfo models.MsgInfo, topic string) error {
	bytes, err := json.Marshal(msgInfo)
	if err != nil {
		return err
	}
	return n.conn.Publish(topic, bytes)
}
