package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
	"time"
)

type BudgetRepository struct {
	storage storage.BudgetStorageInterface
}

func NewBudgetRepository(storage storage.BudgetStorageInterface) *BudgetRepository { //конструктор
	return &BudgetRepository{
		storage: storage,
	}
}

func (b *BudgetRepository) CreateBudget(ctx context.Context, budget models.Budget) (models.Budget, error) {
	query := `INSERT INTO budgets (user_id, category_id, amount, spent_amount, period, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result, err := b.storage.CreateBudget(ctx, query, budget)
	if err != nil {
		return models.Budget{}, err
	}
	return result, nil
}

func (b *BudgetRepository) GetUserBudgets(ctx context.Context, category_id int, userID uint) ([]models.Budget, error) {
	query := `
		SELECT id, user_id, category_id, amount, spent_amount, period, start_date, end_date
		FROM budgets
		WHERE user_id = $1 AND ($2 = 0 OR category_id = $2)
		ORDER BY start_date DESC
	`
	result, err := b.storage.GetUserBudgets(ctx, query, category_id, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BudgetRepository) GetBudgetByID(ctx context.Context, userID uint, category_id int, budget_id int) (models.Budget, error) {
	query := `
	SELECT id, user_id, category_id, amount, spent_amount, period, start_date, end_date
	FROM budgets WHERE id = $1 AND user_id = $2 AND ($3 = 0 OR category_id = $3)`
	result, err := b.storage.GetBudgetByID(ctx, query, userID, category_id, budget_id)
	if err != nil {
		return models.Budget{}, err
	}
	return result, nil
}

func (b *BudgetRepository) DeleteBudgetsInCategory(ctx context.Context, userID uint, categoryID int) error {
	query := `DELETE FROM budgets WHERE user_id = $1 AND category_id = $2`
	err := b.storage.DeleteBudgetsInCategory(ctx, query, userID, categoryID)
	if err != nil {
		return err
	}
	return nil
}

func (b *BudgetRepository) DeleteBudget(ctx context.Context, userID uint, category_id int, budget_id int) error {
	query := `DELETE FROM budgets WHERE id = $1 AND user_id = $2 AND ($3 = 0 OR category_id = $3)`
	err := b.storage.DeleteBudget(ctx, query, userID, category_id, budget_id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSpentAmount обновляет потраченную сумму для бюджета
func (b *BudgetRepository) UpdateSpentAmount(ctx context.Context, category_id int, budgetID uint, spentAmount float64) error {
	query := `UPDATE budgets SET spent_amount = $1 WHERE id = $2 AND ($3 = 0 OR category_id = $3)`
	err := b.storage.UpdateSpentAmount(ctx, query, category_id, budgetID, spentAmount)
	if err != nil {
		return err
	}
	return nil
}

func (b *BudgetRepository) GetActiveBudgetsByCategoryAndDate(ctx context.Context, userID uint, categoryID int, date time.Time) ([]models.Budget, error) {
	query := `
		SELECT id, user_id, category_id, amount, spent_amount, period, start_date, end_date
		FROM budgets
		WHERE user_id = $1 AND category_id = $2 
		  AND ($3 BETWEEN start_date AND end_date OR (start_date IS NULL AND end_date IS NULL))
		ORDER BY start_date DESC
	`
	result, err := b.storage.GetActiveBudgetsByCategoryAndDate(ctx, query, userID, categoryID, date)
	if err != nil {
		return nil, err
	}
	return result, nil
}
