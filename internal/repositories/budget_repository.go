package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type BudgetRepository struct {
	pool *pgxpool.Pool
}

func NewBudgetRepository(pool *pgxpool.Pool) *BudgetRepository { //конструктор
	return &BudgetRepository{
		pool: pool,
	}
}
