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

func (b *BudgetRepository) CreateBudget(ctx context.Context, budget *models.Budget) error {

}

func (b *BudgetRepository) GetBudgetByID(ctx context.Context, id uint) (*models.Budget, error) {

}

func (b *BudgetRepository) GetBudgetByUserID(ctx context.Context, userID uint) ([]*models.Budget, error) {

}

func (b *BudgetRepository) DeleteBudget(ctx context.Context, id uint) error {

}
