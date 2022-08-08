package postgres

import (
	"database/sql"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

type AuthorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (r *AuthorRepo) CreateAuthor(a *models.Author) (*models.Author, error) {

	query := `INSERT INTO authors (id, first_name, last_name, photo_link) 
	VALUES($1, $2, $3, $4) RETURNING *`

	row := r.db.QueryRow(query, a.Id, a.FirstName, a.LastName, a.PhotoLink)

	var author models.Author
	err := row.Scan(
		&author.Id,
		&author.FirstName,
		&author.LastName,
		&author.PhotoLink,
	)
	if err != nil {
		return nil, err
	}

	return &author, nil
}
