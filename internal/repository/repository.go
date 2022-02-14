package repository

import (
	"github.com/Demoss/books/internal/domain"
	"github.com/jmoiron/sqlx"
)

const (
	users = "users"
)

type Books interface {
}

type Authorization interface {
	CreateUser(user domain.User) error
	GetUser(username, password string) (domain.User, error)
}
type Repository struct {
	Books
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books:         NewBooksPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
