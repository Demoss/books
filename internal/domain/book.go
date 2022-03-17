package domain

type Book struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	Impression string `json:"impression"`
}
