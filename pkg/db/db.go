package db

import (
	"MessagioTestTask/pkg/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	User     string `yaml:"user" env:"USER" env-default:"postgres"`
	Password string `yaml:"password" env:"PASSWORD" env-default:"postgres"`
	DbName   string `yaml:"dbName" env:"DB_NAME" env-default:"testtask"`
}

type DB struct {
	cfg *Config
	db  *pgx.Conn
}

func New(cfg *Config) (*DB, error) {
	d := &DB{
		cfg: cfg,
	}

	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName))
	if err != nil {
		return nil, err
	}
	d.db = db

	return d, nil
}

func (d *DB) Close() error {
	return d.db.Close(context.Background())
}

func (d *DB) AddMessage(message models.Message) (int64, error) {
	var id int64

	err := d.db.QueryRow(context.Background(),
		`INSERT INTO public.messages (title, data)
			 VALUES ($1, $2) RETURNING id`, message.Title, message.Data).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
