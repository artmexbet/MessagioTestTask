package service

import (
	"MessagioTestTask/pkg/models"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log/slog"
)

// IDatabase ...
type IDatabase interface {
	AddMessage(models.Message) (int64, error)
}

type IBroker interface {
	SendMsgInfo(msgInfo models.MsgInfo, topic string) error
}

type Config struct {
}

type Service struct {
	cfg    *Config
	db     IDatabase
	broker IBroker
}

func New(cfg *Config, db IDatabase, broker IBroker) (*Service, error) {
	s := &Service{
		cfg:    cfg,
		db:     db,
		broker: broker,
	}

	return s, nil
}

func (s *Service) HandleMessage() nats.MsgHandler {
	return func(m *nats.Msg) {
		var msg models.Message
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			slog.Error("Failed to unmarshal message", slog.String("error", err.Error()))
			return
		}

		id, err := s.db.AddMessage(msg)
		if err != nil {
			slog.Error("Failed to add message to db",
				slog.String("error", err.Error()),
				slog.String("msg_id", msg.Id))
		} else {
			slog.Info("Message added to db", slog.Int64("id", id))

			msgInfo := models.MsgInfo{
				Id: id,
			}

			err := s.broker.SendMsgInfo(msgInfo, "process_message")
			if err != nil {
				slog.Error("Failed to send message to broker",
					slog.String("error", err.Error()),
					slog.String("msg_id", msg.Id))
			} else {
				slog.Info("Message sent to broker", slog.String("msg_id", msg.Id))
			}
		}
	}
}
