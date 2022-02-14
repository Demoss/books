package domain

type Book struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title"`
	Author     `json:"author" db:"author"`
	Impression string `json:"impression"`
}
