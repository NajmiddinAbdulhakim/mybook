package service

import (
	"context"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/NajmiddinAbdulhakim/mybook/storage"
)

type Service struct {
	repo storage.IStorage
}

func NewService(repo storage.IStorage) Service {
	return Service{repo: repo}
}

func (s *Service) CreateBook(ctx context.Context, b *models.Book) (*models.Book, error) {
	return nil, nil
}
