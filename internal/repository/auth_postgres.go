package repository

import (
	"fmt"
	"github.com/Demoss/books/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password, is_author) values ($1, $2, $3)", users)
	r.db.QueryRow(query, user.Username, user.Password, user.IsAuthor)
	return nil
}

func (r *AuthPostgres) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", users)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
