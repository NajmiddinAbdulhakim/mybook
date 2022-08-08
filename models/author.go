package models

type Author struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhotoLink string `json:"photo_link"`
}
