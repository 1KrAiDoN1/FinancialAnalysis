package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/repositories"
)

type BudgetService struct {
	repo repositories.BudgetRepositoryInterface
}

func NewBudgetService(repo repositories.BudgetRepositoryInterface) *BudgetService {
	return &BudgetService{
		repo: repo,
	}
}

func (b *BudgetService) CreateBudget(ctx context.Context, userID uint, req dto.CreateBudgetRequest) (*dto.BudgetResponse, error) {

}

func (b *BudgetService) GetUserBudgets(ctx context.Context, userID uint) ([]*dto.BudgetResponse, error) {

}

func (b *BudgetService) UpdateBudget(ctx context.Context, userID uint, budgetID uint, req dto.UpdateBudgetRequest) error {

}

func (b *BudgetService) DeleteBudget(ctx context.Context, userID uint, budgetID uint) error {

}
