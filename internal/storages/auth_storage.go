package storage

import "github.com/jackc/pgx/v5/pgxpool"

type AuthStorage struct {
	pool *pgxpool.Pool
}

func NewAuthStorage(pool *pgxpool.Pool) *AuthStorage {
	return &AuthStorage{
		pool: pool,
	}
}
