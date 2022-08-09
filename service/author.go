package service

import (
	"context"
	"log"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateAuthor(ctx context.Context, req *models.Author) (*models.Author, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	req.Id = id.String()

	res, err := s.repo.Author().CreateAuthor(req)
	if err != nil {
		log.Fatal(`filed while inserting author: `, err)
		return nil, status.Error(codes.Internal, `filed while inserting author`)
	}

	return res, nil
}

func (s *Service) UpdateAuthor(ctx context.Context, req *models.Author) (*models.Author, error) {
	res, err := s.repo.Author().UpdateAuthor(req)
	if err != nil {
		log.Fatal(`filed while updating author: `, err)
		return nil, status.Error(codes.Internal, `filed while updating author`)
	}

	return res, nil
}

func (s *Service) GetAuthor(ctx context.Context, req string) (*models.Author, error) {
	res, err := s.repo.Author().GetAuthor(req)
	if err != nil {
		log.Fatal(`filed while getting author: `, err)
		return nil, status.Error(codes.Internal, `filed while getting author`)
	}
	return res, nil
}

func (s *Service) DeleteAuthor(ctx context.Context, req string) error {
	err := s.repo.Author().DeleteAuthor(req)
	if err != nil {
		log.Fatal(`filed while deleting author: `, err)
		return status.Error(codes.Internal, `filed whele deleting author`)
	}
	
	return nil
}
