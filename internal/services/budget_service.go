package services

import (
	"context"
	"errors"
	"finance/internal/dto"
	"finance/internal/models"
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

func (b *BudgetService) CreateBudget(ctx context.Context, userID uint, req dto.CreateBudgetRequest) (dto.BudgetResponse, error) {
	req_budget := models.Budget{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Period:     req.Period,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}
	res_budget, err := b.repo.CreateBudget(ctx, req_budget)
	if err != nil {
		return dto.BudgetResponse{}, err
	}
	return dto.BudgetResponse{
		ID:         res_budget.ID,
		CategoryID: res_budget.CategoryID,
		Amount:     res_budget.Amount,
		Period:     res_budget.Period,
		StartDate:  res_budget.StartDate,
		EndDate:    res_budget.EndDate,
		CreatedAt:  res_budget.CreatedAt,
	}, nil

}

func (b *BudgetService) GetUserBudgets(ctx context.Context, userID uint) ([]dto.BudgetResponse, error) {
	budgets, err := b.repo.GetUserBudgets(ctx, userID)
	if err != nil {
		return nil, err
	}
	budgetResponses := make([]dto.BudgetResponse, len(budgets))

	// Преобразуем каждый элемент из models.Budget в dto.BudgetResponse
	for i, budget := range budgets {
		budgetResponses[i] = dto.BudgetResponse{
			ID:         budget.ID,
			CategoryID: budget.CategoryID,
			Amount:     budget.Amount,
			Period:     budget.Period,
			StartDate:  budget.StartDate,
			EndDate:    budget.EndDate,
			CreatedAt:  budget.CreatedAt,
		}
	}
	return budgetResponses, nil
}

func (b *BudgetService) UpdateBudget(ctx context.Context, userID uint, budgetID int, req dto.UpdateBudgetRequest) error {
	existingBudget, err := b.repo.GetBudgetByID(ctx, userID, budgetID)
	if err != nil {
		return err
	}

	// Проверяем права доступа
	if existingBudget.UserID != userID {
		return errors.New("budget not found or access denied")
	}

	// Создаем объект для обновления
	updateData := models.Budget{
		ID:         uint(budgetID),
		UserID:     userID,
		CategoryID: req.CategoryID,
		Amount:     *req.Amount,
		Period:     *req.Period,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}

	// Обновляем через репозиторий
	return b.repo.UpdateBudget(ctx, updateData)
}

func (b *BudgetService) DeleteBudget(ctx context.Context, userID uint, budgetID int) error {
	return b.repo.DeleteBudget(ctx, userID, uint(budgetID))
}
