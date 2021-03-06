package service

import (
	"github.com/Demoss/books/internal/domain"
	"github.com/Demoss/books/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Books interface {
	AddBook(book domain.Book) error
	GetAuthorsBooks(author domain.Author) ([]domain.Book, error)
	DeleteBook(book domain.Book) error
}

type Service struct {
	Authorization
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books:         NewBooksService(repos.Books),
		Authorization: NewAuthService(repos.Authorization),
	}
}
