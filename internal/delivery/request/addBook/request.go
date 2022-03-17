package addBook

import "github.com/Demoss/books/internal/domain"

type Request struct {
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	Impression string `json:"impression"`
}

func MapToDomain(request Request) domain.Book {
	return domain.Book{
		Title:      request.Title,
		AuthorID:   request.AuthorID,
		Impression: request.Impression,
	}
}
