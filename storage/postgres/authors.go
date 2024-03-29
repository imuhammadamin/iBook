package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"iBook/api/models"
)

type authorsRepo struct {
	db *pgxpool.Pool
}

func NewAuthorsRepo(db *pgxpool.Pool) authorsRepo {
	return authorsRepo{db: db}
}

func (a *authorsRepo) Create(author models.Author) error {
	query := `INSERT INTO authors (firstname, lastname, age) VALUES ($1, $2, $3)`

	ctag, err := a.db.Exec(context.Background(), query, author.Firstname, author.Lastname, author.Age)
	if err != nil {
		return err
	}

	fmt.Println(ctag.String())
	fmt.Println(ctag.RowsAffected())

	return nil
}

func (a *authorsRepo) Update(author models.Author) error {
	query := `UPDATE authors SET firstname=$1, lastname=$2, age=$3 WHERE id=$4`

	_, err := a.db.Exec(context.Background(), query, author.Firstname, author.Lastname, author.Age, author.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *authorsRepo) GetById(id int) (models.Author, error) {
	author := models.Author{}

	query := `SELECT * FROM authors WHERE id=$1`

	rows, err := a.db.Query(context.Background(), query, id)
	if err != nil {
		return author, err
	}

	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Firstname, &author.Lastname, &author.Age)
		if err != nil {
			return author, err
		}
	}

	return author, nil
}

func (a *authorsRepo) GetAll(paginationMetadata models.PaginationMetadata) ([]models.Author, error) {
	var authors []models.Author
	offset := (paginationMetadata.Page - 1) * paginationMetadata.Limit

	query := `SELECT * FROM authors LIMIT $1 OFFSET $2`

	rows, err := a.db.Query(context.Background(), query, paginationMetadata.Limit, offset)
	if err != nil {
		return authors, err
	}

	for rows.Next() {
		author := models.Author{}

		err := rows.Scan(&author.ID, &author.Firstname, &author.Lastname, &author.Age)
		if err != nil {
			return authors, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a *authorsRepo) Delete(id int) error {
	query := `DELETE FROM authors WHERE id=$1`

	_, err := a.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
