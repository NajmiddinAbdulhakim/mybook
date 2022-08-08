package service

import (
	"context"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type Repository interface {
	CreateAuthor(ctx context.Context, req *models.Author) (*models.Author, error)
	CreateUser(ctx context.Context, req *models.User) (*models.User, error)
	CreateBook(ctx context.Context, req *models.Book) (*models.Book, error)
}


