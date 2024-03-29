package storage

import "iBook/api/models"

type IStorage interface {
	CloseDB()
	Author() IAuthorStorage
	Book() IBookStorage
}

type IAuthorStorage interface {
	Create(models.Author) error
	GetById(int) (models.Author, error)
	GetAll(models.PaginationMetadata) ([]models.Author, error)
	Update(models.Author) error
	Delete(int) error
}

type IBookStorage interface {
	Create(models.Book) error
	GetById(int) (models.Book, error)
	GetAll(models.PaginationMetadata) ([]models.Book, error)
	Update(models.Book) error
	Delete(int) error
}
