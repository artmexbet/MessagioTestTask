package service

import (
	"MessagioTestTask/pkg/models"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type IBroker interface {
	SendMessage(message models.Message, topic string) error
}

// Config ...
type Config struct {
}

// Service ...
type Service struct {
	cfg       *Config
	validator *validator.Validate
	queue     IBroker
}

// New ...
func New(cfg *Config, queue IBroker) (*Service, error) {
	s := &Service{
		cfg:       cfg,
		queue:     queue,
		validator: validator.New(),
	}

	return s, nil
}

// PostMessage validate and send message to kafka from http request
func (s *Service) PostMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var msgJson models.AddMessageJSON
		if err := json.NewDecoder(r.Body).Decode(&msgJson); err != nil {
			slog.Error("Failed to decode message", slog.String("error", err.Error()))
			http.Error(w, "Failed to decode message", http.StatusBadRequest)
			return
		}

		if err := s.validator.Struct(msgJson); err != nil {
			slog.Error("Failed to validate message", slog.String("error", err.Error()))
			http.Error(w, "Failed to validate message", http.StatusBadRequest)
			return
		}

		var msg models.Message
		msg.Id = middleware.GetReqID(r.Context())
		msg.Data = msgJson.Data
		msg.Title = msgJson.Title

		err := s.queue.SendMessage(msg, "add_message")

		if err != nil {
			slog.Error("Failed to write message", slog.String("error", err.Error()))
			http.Error(w, "Failed to write message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Message sent"))
	}
}
