package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	booksTable = "books"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSlMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.DBName,
			cfg.Password,
			cfg.SSlMode,
		),
	)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
