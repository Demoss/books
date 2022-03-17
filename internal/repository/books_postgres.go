package repository

import (
	"fmt"
	"github.com/Demoss/books/internal/domain"
	"github.com/jmoiron/sqlx"
)

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBooksPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) AddBook(book domain.Book) error {
	query := fmt.Sprintf("INSERT INTO %s (title, author_id,impression) values ($1, $2, $3)", books)
	if err := r.db.QueryRow(query, book.Title, book.AuthorID, book.Impression); err != nil {
		return err.Err()
	}
	return nil
}

func (r *BooksPostgres) GetAuthorsBooks(author domain.Author) ([]domain.Book, error) {
	var res []domain.Book
	query := fmt.Sprintf("SELECT b.title,b.impression FROM %s b INNER JOIN %s a on b.author_id = a.id WHERE a.name = $1 AND a.surname = $2", books, authors)
	err := r.db.Select(&res, query, author.Name, author.Surname)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *BooksPostgres) DeleteBook(book domain.Book) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE title = $1 AND author_id = $2", books)
	_, err := r.db.Exec(query, book.Title, book.AuthorID)
	if err != nil {
		return err
	}
	return nil
}
