package sign_in

import "github.com/Demoss/books/internal/domain"

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func MapToDomain(request Request) domain.User {
	return domain.User{Username: request.Username, Password: request.Password}
}
