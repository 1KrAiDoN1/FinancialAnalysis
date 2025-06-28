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
	return models.Budget{}, nil

}

func (b *BudgetRepository) GetUserBudgets(ctx context.Context, category_id int, userID uint) ([]models.Budget, error) {
	return []models.Budget{}, nil
}

func (b *BudgetRepository) GetBudgetByID(ctx context.Context, userID uint, category_id int, budget_id int) (models.Budget, error) {
	return models.Budget{}, nil
}

func (b *BudgetRepository) DeleteBudgetsInCategory(ctx context.Context, userID uint, categoryID int) error {
	return nil
}

//	func (b *BudgetRepository) UpdateBudget(ctx context.Context, budget models.Budget) error {
//		return nil
//	}
func (b *BudgetRepository) DeleteBudget(ctx context.Context, userID uint, category_id int, budget_id int) error {
	return nil
}

// UpdateSpentAmount обновляет потраченную сумму для бюджета
func (b *BudgetRepository) UpdateSpentAmount(ctx context.Context, category_id int, budgetID uint, spentAmount float64) error {
	return nil
}

func (b *BudgetRepository) GetActiveBudgetsByCategoryAndDate(ctx context.Context, userID uint, categoryID int, date time.Time) ([]models.Budget, error) {
	return nil, nil
}
