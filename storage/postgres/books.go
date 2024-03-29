package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"iBook/api/models"
)

type booksRepo struct {
	db *pgxpool.Pool
}

func NewBooksRepo(db *pgxpool.Pool) booksRepo {
	return booksRepo{db: db}
}

func (b *booksRepo) Create(book models.Book) error {
	query := `INSERT INTO books (name, genre, description, release_year, author_id) VALUES ($1, $2, $3, $4, $5)`

	_, err := b.db.Exec(context.Background(), query, book.Name, book.Genre, book.Description, book.ReleaseYear, book.AuthorId)
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) Update(book models.Book) error {
	query := `UPDATE books SET name=$1, genre=$2, description=$3, release_year=$4 WHERE id=$5`

	_, err := b.db.Exec(context.Background(), query, book.Name, book.Genre, book.Description, book.ReleaseYear, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) GetById(id int) (models.Book, error) {
	var book models.Book

	query := `SELECT * FROM books b INNER JOIN authors a ON b.author_id = a.id WHERE b.id=$1`
	rows, err := b.db.Query(context.Background(), query, id)
	if err != nil {
		return book, err
	}

	for rows.Next() {
		book = models.Book{}
		author := models.Author{}

		err = rows.Scan(&book.ID, &book.Name, &book.Genre, &book.Description, &book.ReleaseYear,
			&book.AuthorId, &author.ID, &author.Firstname, &author.Lastname, &author.Age)

		book.Author = author
	}

	return book, nil
}

func (b *booksRepo) GetAll(paginationMetadata models.PaginationMetadata) ([]models.Book, error) {
	var books []models.Book
	offset := (paginationMetadata.Page - 1) * paginationMetadata.Limit

	query := `SELECT * FROM books LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(context.Background(), query, paginationMetadata.Limit, offset)
	if err != nil {
		return books, err
	}

	for rows.Next() {
		book := models.Book{}

		err := rows.Scan(&book.ID, &book.Name, &book.Genre, &book.Description, &book.ReleaseYear, &book.AuthorId)
		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b *booksRepo) Delete(id int) error {
	query := `DELETE FROM books WHERE id=$1`

	_, err := b.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
