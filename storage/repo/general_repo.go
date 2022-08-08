package repo

import (
	"context"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type UserStorage interface {
	CreateUser(ctx context.Context, u *models.User) (*models.User, error)
}
type AuthorStorage interface {
	CreateAuthor(ctx context.Context, a *models.Author) (*models.Author, error)

}
type BookStorage interface {
	CreateBook(ctx context.Context, b *models.Book) (*models.Book, error)
}
