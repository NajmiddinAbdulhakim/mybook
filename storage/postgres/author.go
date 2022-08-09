package postgres

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"

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

func (r *AuthorRepo) UpdateAuthor(a *models.Author) (*models.Author, error) {
	query := `UPDATE authors 
	SET first_name = $1, last_name = $2, photo_link = $3
	WHERE id = $4 RETURNING *`

	var author models.Author
	row := r.db.QueryRow(query, a.FirstName, a.LastName, a.PhotoLink, a.Id)
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

func (r *AuthorRepo) GetAuthor(id string) (*models.Author, error) {
	query := `SELECT * FROM authors
	WHERE id = $1`

	var author models.Author
	row := r.db.QueryRow(query, id)
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

func (r *AuthorRepo) DeleteAuthor(id string) error {
	query := `DELETE FROM authors 
	WHERE id = $1`
	row,err := r.db.Exec(query, id)
	con,_ := row.RowsAffected()
	if err != nil || 0 == con {
		return errors.New(`row 0`)
	}
	return nil
}

func (r *AuthorRepo) cleanTable() error {
	query := `DELETE FROM authors`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
