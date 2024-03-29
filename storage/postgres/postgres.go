package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"iBook/config"
	"iBook/storage"
	"time"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf("host=%s port=%v database=%s user=%s password=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase, cfg.PostgresUser, cfg.PostgresPassword)

	pgPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	pgPoolConfig.MaxConns = 100
	pgPoolConfig.MaxConnLifetime = time.Minute * 2

	newPool, err := pgxpool.NewWithConfig(context.Background(), pgPoolConfig)
	if err != nil {
		return nil, err
	}

	return Store{Pool: newPool}, nil
}

func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) Author() storage.IAuthorStorage {
	newAuthor := NewAuthorsRepo(s.Pool)
	return &newAuthor
}

func (s Store) Book() storage.IBookStorage {
	newBook := NewBooksRepo(s.Pool)
	return &newBook
}
