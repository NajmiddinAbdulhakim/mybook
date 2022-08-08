package postgres

import (
	"database/sql"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(u *models.User) (*models.User, error) {
	return nil, nil

}
