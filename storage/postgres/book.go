package postgres

import (
	"context"
	"database/sql"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) CreateBook(ctx context.Context, b *models.Book) (*models.Book, error) {
	return nil, nil

}
