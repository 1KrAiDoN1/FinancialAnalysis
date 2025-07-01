package storage

import (
	"context"
	"finance/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseStorage struct {
	pool *pgxpool.Pool
}

func NewExpenseStorage(pool *pgxpool.Pool) *ExpenseStorage {
	return &ExpenseStorage{
		pool: pool,
	}
}

func (s *ExpenseStorage) CreateExpense(ctx context.Context, query string, expense models.Expense) (models.Expense, error) {
	var new_expense models.Expense
	err := s.pool.QueryRow(ctx, query, expense.UserID, expense.CategoryID, expense.Amount, expense.Description, expense.Date, expense.CreatedAt).Scan(&new_expense.ID, &new_expense.CategoryID, &new_expense.Amount, &new_expense.Description, &new_expense.Date, &new_expense.CreatedAt)
	if err != nil {
		return models.Expense{}, err
	}
	return expense, nil
}
