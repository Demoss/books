package service

import (
	"github.com/Demoss/books/internal/repository"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}
