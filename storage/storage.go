package storage

import (
	"database/sql"

	"github.com/NajmiddinAbdulhakim/mybook/storage/postgres"
	"github.com/NajmiddinAbdulhakim/mybook/storage/repo"
)

type IStorage interface {
	Book() repo.BookStorage
	Author() repo.AuthorStorage
	User() repo.UserStorage
}

type storagePg struct {
	db         *sql.DB
	bookRepo   repo.BookStorage
	authorRepo repo.AuthorStorage
	userRepo   repo.UserStorage
}

func NewStoragePg(db *sql.DB) *storagePg {
	return &storagePg{
		db:         db,
		bookRepo:   postgres.NewBookRepo(db),
		authorRepo: postgres.NewAuthorRepo(db),
		userRepo:   postgres.NewUserRepo(db),
	}
}

func (s *storagePg) Book() repo.BookStorage {
	return s.bookRepo
}

func (s *storagePg) Author() repo.AuthorStorage {
	return s.authorRepo
}

func (s *storagePg) User() repo.UserStorage {
	return s.userRepo
}
