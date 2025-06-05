package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type ExpenseRepository struct {
	pool *pgxpool.Pool
}

func NewExpenseRepository(pool *pgxpool.Pool) *ExpenseRepository { //конструктор
	return &ExpenseRepository{
		pool: pool,
	}
}
