package storage

import "github.com/jackc/pgx/v5/pgxpool"

type ExpenseStorage struct {
	pool *pgxpool.Pool
}

func NewExpenseStorage(pool *pgxpool.Pool) *ExpenseStorage {
	return &ExpenseStorage{
		pool: pool,
	}
}
