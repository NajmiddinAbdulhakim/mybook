package postgres

import (
	"database/sql"
	"context"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type AuthorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (r *AuthorRepo) CreateAuthor(ctx context.Context, a *models.Author) (*models.Author, error) {
	return nil, nil
}