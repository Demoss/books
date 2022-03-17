package service

import (
	"github.com/Demoss/books/internal/domain"
	"github.com/Demoss/books/internal/repository"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) AddBook(book domain.Book) error {
	return s.repo.AddBook(book)
}

func (s *BooksService) GetAuthorsBooks(author domain.Author) ([]domain.Book, error) {
	return s.repo.GetAuthorsBooks(author)
}

func (s *BooksService) DeleteBook(book domain.Book) error {
	return s.repo.DeleteBook(book)
}
