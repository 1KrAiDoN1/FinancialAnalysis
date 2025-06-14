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

func (e *ExpenseRepository) CreateExpense(ctx context.Context, expense *models.Expense) error {

}

func (e *ExpenseRepository) GetExpenseByID(ctx context.Context, id uint) (*models.Expense, error) {

}

func (e *ExpenseRepository) GetExpenseByUserID(ctx context.Context, userID uint, expense *models.Expense) ([]*models.Expense, error) {

}

func (e *ExpenseRepository) DeleteExpense(ctx context.Context, id uint) error {

}

func (e *ExpenseRepository) GetExpensesByDateRange(ctx context.Context, userID uint, startDate, endDate time.Time) ([]*models.Expense, error) {

}

func (e *ExpenseRepository) GetExpensesByCategory(ctx context.Context, userID uint, categoryID uint, limit, offset int) ([]*models.Expense, error) {

}

func (e *ExpenseRepository) GetAverageExpenseAmount(ctx context.Context, userID uint) (float64, error) {

}

func (e *ExpenseRepository) GetLargestExpense(ctx context.Context, userID uint) (*models.Expense, error) {

}

func (e *ExpenseRepository) GetSmallestExpense(ctx context.Context, userID uint) (*models.Expense, error) {

}
