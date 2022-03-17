package deleteBook

import "github.com/Demoss/books/internal/domain"

type Request struct {
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}

func MapToCommand(request Request) domain.Book {
	return domain.Book{
		Title:    request.Title,
		AuthorID: request.AuthorID,
	}
}
