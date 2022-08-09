package service

import (
	"context"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/NajmiddinAbdulhakim/mybook/storage"
)

type Repository interface {
	CreateAuthor(ctx context.Context, req *models.Author) (*models.Author, error)
	UpdateAuthor(ctx context.Context, req *models.Author) (*models.Author, error)
	GetAuthor(ctx context.Context, req string) (*models.Author, error)
	DeleteAuthor(ctx context.Context, req string) error

	CreateUser(ctx context.Context, req *models.User) (*models.User, error)

	CreateBook(ctx context.Context, req *models.Book) (*models.Book, error)
}

type Service struct {
	repo storage.IStorage
}

func NewService(repo storage.IStorage) Service {
	return Service{repo: repo}
}
