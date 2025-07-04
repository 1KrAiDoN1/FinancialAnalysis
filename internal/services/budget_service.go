package services

import (
	"context"
	//"errors"
	"finance/internal/dto"
	"finance/internal/models"
	"finance/internal/repositories"
	"finance/pkg"
	"time"
)

type BudgetService struct {
	repo         repositories.BudgetRepositoryInterface
	expense_repo repositories.ExpenseRepositoryInterface
}

func NewBudgetService(repo repositories.BudgetRepositoryInterface, expense_repo repositories.ExpenseRepositoryInterface) *BudgetService {
	return &BudgetService{
		repo:         repo,
		expense_repo: expense_repo,
	}
}

func (b *BudgetService) CreateBudget(ctx context.Context, userID uint, category_id int, req dto.CreateBudgetRequest) (dto.BudgetResponse, error) {
	start_date := time.Now()
	endDate, err := pkg.AddPeriodToDate(start_date, req.Period)
	if err != nil {
		return dto.BudgetResponse{}, err
	}
	req_budget := models.Budget{
		UserID:      req.UserID,
		CategoryID:  uint(category_id),
		Amount:      req.Amount,
		SpentAmount: 0,
		Period:      req.Period,
		StartDate:   start_date,
		EndDate:     endDate,
	}

	res_budget, err := b.repo.CreateBudget(ctx, req_budget)
	if err != nil {
		return dto.BudgetResponse{}, err
	}

	// Пересчитываем потраченную сумму для нового бюджета
	err = b.recalculateBudgetSpentAmount(ctx, &res_budget)
	if err != nil {
		return dto.BudgetResponse{}, err
	}

	return dto.BudgetResponse{
		ID:              res_budget.ID,
		CategoryID:      res_budget.CategoryID,
		Amount:          res_budget.Amount,
		SpentAmount:     res_budget.SpentAmount,
		RemainingAmount: res_budget.Amount - res_budget.SpentAmount,
		Period:          res_budget.Period,
		StartDate:       res_budget.StartDate,
		EndDate:         res_budget.EndDate,
	}, nil

}

func (b *BudgetService) GetUserBudgets(ctx context.Context, userID uint, category_id int) ([]dto.BudgetResponse, error) {
	budgets, err := b.repo.GetUserBudgets(ctx, category_id, userID)
	if err != nil {
		return nil, err
	}
	budgetResponses := make([]dto.BudgetResponse, len(budgets))

	// Преобразуем каждый элемент из models.Budget в dto.BudgetResponse
	for i, budget := range budgets {
		budgetResponses[i] = dto.BudgetResponse{
			ID:          budget.ID,
			CategoryID:  uint(category_id),
			Amount:      budget.Amount,
			SpentAmount: budget.SpentAmount,
			Period:      budget.Period,
			StartDate:   budget.StartDate,
			EndDate:     budget.EndDate,
		}
	}
	return budgetResponses, nil
}

func (b *BudgetService) DeleteBudget(ctx context.Context, userID uint, category_id int, budgetID int) error {
	return b.repo.DeleteBudget(ctx, userID, category_id, budgetID)
}

// recalculateBudgetSpentAmount пересчитывает потраченную сумму для бюджета
func (b *BudgetService) recalculateBudgetSpentAmount(ctx context.Context, budget *models.Budget) error {
	expenses, err := b.expense_repo.GetExpensesByCategoryAndPeriod(ctx, budget.UserID, int(budget.CategoryID), budget.StartDate, budget.EndDate)
	if err != nil {
		return err
	}

	var totalSpent float64
	for _, expense := range expenses {
		totalSpent += expense.Amount
	}
	budget.SpentAmount = totalSpent
	return nil
}
