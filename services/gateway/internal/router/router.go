package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

// Config for router
type Config struct {
	Host string `yaml:"host" env:"HOST" env-default:""`
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
}

// IService interface for routes
type IService interface {
	PostMessage() http.HandlerFunc
}

// Router provides application to handle some requests
type Router struct {
	cfg    *Config
	svc    IService
	router *chi.Mux
}

// New initializes function for Router struct
func New(cfg *Config, svc IService) *Router {
	r := &Router{
		cfg:    cfg,
		svc:    svc,
		router: chi.NewRouter(),
	}

	r.router.Use(middleware.Recoverer)
	r.router.Use(middleware.Logger)
	r.router.Use(middleware.RequestID)
	r.router.Use(middleware.AllowContentType("application/json"))
	r.router.Use(middleware.Timeout(10 * time.Second))
	r.router.Use(middleware.RequestSize(5 << 20))
	r.router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// Allow All origins because there are no info about allowed origins in the task
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.router.Post("/api/message", r.svc.PostMessage())

	return r
}

func (r *Router) Run() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", r.cfg.Host, r.cfg.Port), r.router)
}
