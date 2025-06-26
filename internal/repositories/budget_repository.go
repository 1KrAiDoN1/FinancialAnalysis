package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
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

func (b *BudgetRepository) GetUserBudgets(ctx context.Context, userID uint) ([]models.Budget, error) {
	return []models.Budget{}, nil
}

func (b *BudgetRepository) GetBudgetByID(ctx context.Context, userID uint, budget_id int) (models.Budget, error) {
	return models.Budget{}, nil
}

func (b *BudgetRepository) GetBudgetByUserID(ctx context.Context, userID uint) ([]models.Budget, error) {
	return nil, nil
}

func (b *BudgetRepository) UpdateBudget(ctx context.Context, budget models.Budget) error {
	return nil
}
func (b *BudgetRepository) DeleteBudget(ctx context.Context, userID uint, id uint) error {
	return nil
}
