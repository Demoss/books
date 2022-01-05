package repository

import "github.com/jmoiron/sqlx"

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBooksPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}
