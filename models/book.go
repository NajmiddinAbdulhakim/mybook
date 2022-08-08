package models



type Book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	PhotoLink string `json:"photo_link"`
	AuthorId string `json:"author_id"`
}