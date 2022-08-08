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

// CreateUser creates book
// @Summary Create book summary
// @Description This api is using create new book
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param book body models.User true "book body"
// @Router /book/create [post]
func (s *Service) CreateBook(ctx context.Context, b *models.Book) (*models.Book, error) {
	return nil, nil
}

// CreateUser creates author
// @Summary Create author summary
// @Description This api is using create new author
// @Tags Author
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param author body models.User true "author body"
// @Router /author/create [post]
func (s *Service) CreateAuthor(ctx context.Context, b *models.Author) (*models.Author, error) {
	return nil, nil
}
// CreateUser creates user
// @Summary Create user summary
// @Description This api is using create new user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param user body models.User true "user body"
// @Router /user/create [post]
func (s *Service) CreateUser(ctx context.Context, b *models.User) (*models.User, error) {
	return nil, nil
}