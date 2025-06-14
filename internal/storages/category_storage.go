package storage

import "github.com/jackc/pgx/v5/pgxpool"

type CategoryStorage struct {
	pool *pgxpool.Pool
}

func NewCategoryStorage(pool *pgxpool.Pool) *CategoryStorage {
	return &CategoryStorage{
		pool: pool,
	}
}
