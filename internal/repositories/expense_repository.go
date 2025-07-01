package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
	"time"
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
	query := `INSERT INTO expenses (user_id, category_id, amount, description, date, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, category_id, amount, description, date, created_at`
	result, err := e.storage.CreateExpense(ctx, query, expense)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil
}

func (e *ExpenseRepository) GetExpenseByID(ctx context.Context, userID uint, category_id int, id uint) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepository) GetExpensesByUserID(ctx context.Context, category_id int, userID uint) ([]models.Expense, error) {
	return nil, nil
}

func (e *ExpenseRepository) GetExpensesByPeriod(ctx context.Context, userID uint, category_id int, period string) ([]models.Expense, error) {
	return nil, nil
}

func (e *ExpenseRepository) DeleteExpense(ctx context.Context, userID uint, category_id int, id uint) error {
	return nil
}

func (e *ExpenseRepository) DeleteExpensesInCategory(ctx context.Context, userID uint, categoryID int) error {
	return nil
}

func (e *ExpenseRepository) GetExpensesByCategory(ctx context.Context, userID uint, categoryID int, limit, offset int) ([]models.Expense, error) {
	return nil, nil
}

func (e *ExpenseRepository) GetLargestExpenseByPeriod(ctx context.Context, userID uint, category_id int, period string) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepository) GetSmallestExpenseByPeriod(ctx context.Context, userID uint, category_id int, period string) (models.Expense, error) {
	return models.Expense{}, nil
}

func (e *ExpenseRepository) GetExpensesByCategoryAndPeriod(ctx context.Context, userID uint, categoryID int, startDate, endDate time.Time) ([]models.Expense, error) {
	return nil, nil
}
