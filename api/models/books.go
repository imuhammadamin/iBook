package models

type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ReleaseYear int    `json:"release_year"`
	AuthorId    int    `json:"author_id"`
	Author      Author `json:"author"`
}
