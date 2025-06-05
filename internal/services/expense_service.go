package services

import (
	"context"
	"finance/internal/dto"
	repositories "finance/internal/repositories"
)

type ExpenseService struct {
	repo repositories.ExpenseRepositoryInterface
}

func NewExpenseService(repo repositories.ExpenseRepositoryInterface) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) CreateExpense(ctx context.Context, userID uint, req dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
	return &dto.ExpenseResponse{}, nil
}

func (s *ExpenseService) GetUserExpenses(ctx context.Context, userID uint, filter dto.ExpenseFilter) ([]*dto.ExpenseResponse, error) {
	return []*dto.ExpenseResponse{}, nil
}

func (s *ExpenseService) UpdateExpense(ctx context.Context, userID uint, expenseID uint, req dto.UpdateExpenseRequest) error {
	return nil
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, userID uint, expenseID uint) error {
	return nil
}

func (s *ExpenseService) GetExpenseAnalytics(ctx context.Context, userID uint, period string) (*dto.ExpenseAnalytics, error) {
	return &dto.ExpenseAnalytics{}, nil
}
