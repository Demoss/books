package service

import "books/pkg/repository"

type Books interface {
}

type Service struct {
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Books: NewBooksService(repos.Books)}
}
