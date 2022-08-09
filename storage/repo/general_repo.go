package repo

import (
	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type UserStorage interface {
	CreateUser(u *models.User) (*models.User, error)
}
type AuthorStorage interface {
	CreateAuthor(a *models.Author) (*models.Author, error)
	UpdateAuthor(a *models.Author) (*models.Author, error)
	GetAuthor(id string) (*models.Author, error)
	DeleteAuthor(id string) error
}
type BookStorage interface {
	CreateBook(b *models.Book) (*models.Book, error)
}
