package storage

import "github.com/jackc/pgx/v5/pgxpool"

type BudgetStorage struct {
	pool *pgxpool.Pool
}

func NewBudgetStorage(pool *pgxpool.Pool) *BudgetStorage {
	return &BudgetStorage{
		pool: pool,
	}
}
