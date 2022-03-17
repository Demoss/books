package getAuthorsBooks

import "github.com/Demoss/books/internal/domain"

type Request struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Response struct {
	Title      string `json:"title"`
	Impression string `json:"impression"`
}

func MapToQuery(request Request) domain.Author {
	return domain.Author{
		Name:    request.Name,
		Surname: request.Surname,
	}
}

func MapToResponse(books []domain.Book) []Response {
	res := make([]Response, len(books))
	for i, book := range books {
		res[i].Title = book.Title
		res[i].Impression = book.Impression
	}
	return res
}
