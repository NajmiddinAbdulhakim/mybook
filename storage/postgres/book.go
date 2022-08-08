package postgres

import (
	"database/sql"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) CreateBook(b *models.Book) (*models.Book, error) {
	return nil, nil

}
