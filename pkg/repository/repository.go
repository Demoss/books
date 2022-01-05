package repository

import "github.com/jmoiron/sqlx"

type Books interface {
}
type Repository struct {
	Books
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Books: new(BooksPostgres)}
}
