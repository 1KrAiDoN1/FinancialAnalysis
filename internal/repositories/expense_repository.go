package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
)

type ExpenseRepository struct {
	storage storage.ExpenseStorageInterface
}

func NewExpenseRepository(storage storage.ExpenseStorageInterface) *ExpenseRepository { //конструктор
	return &ExpenseRepository{
		storage: storage,
	}
}

func (e *ExpenseRepository) CreateExpense(ctx context.Context, expense models.Expense) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepository) GetExpenseByID(ctx context.Context, userID uint, id uint) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepository) GetExpensesByUserID(ctx context.Context, userID uint) ([]models.Expense, error) {
	return nil, nil
}

func (e *ExpenseRepository) GetExpensesByPeriod(ctx context.Context, userID uint, period string) ([]models.Expense, error) {
	return nil, nil
}

func (e *ExpenseRepository) DeleteExpense(ctx context.Context, userID uint, id uint) error {
	return nil
}

func (e *ExpenseRepository) GetExpensesByCategory(ctx context.Context, userID uint, categoryID uint, limit, offset int) ([]models.Expense, error) {
	return nil, nil
}

func (e *ExpenseRepository) GetLargestExpenseByPeriod(ctx context.Context, userID uint, period string) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepository) GetSmallestExpenseByPeriod(ctx context.Context, userID uint, period string) (models.Expense, error) {
	return models.Expense{}, nil
}
