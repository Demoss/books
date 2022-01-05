package domain

type Books struct {
	Id     int `json:"id" db:"id"`
	Author `json:"author" db:"author"`
	Title  string `json:"title"`
}
type Author struct {
	Name string
	Age  int
}
