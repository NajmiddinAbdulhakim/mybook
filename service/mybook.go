package service

import (
	"context"
	"log"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/NajmiddinAbdulhakim/mybook/storage"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Service) CreateAuthor(ctx context.Context, req *models.Author) (*models.Author, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	req.Id = id.String()

	author, err := s.repo.Author().CreateAuthor(req)
	if err != nil {
		log.Fatal(`filed while inserting author: `, err)
		return nil, status.Error(codes.Internal, `filed while inserting author`)
	}

	return author, nil
}

func (s *Service) CreateUser(ctx context.Context, b *models.User) (*models.User, error) {
	return nil, nil
}
