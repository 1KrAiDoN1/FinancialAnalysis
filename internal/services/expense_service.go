package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/models"
	repositories "finance/internal/repositories"
	"time"
)

type ExpenseService struct {
	repo repositories.ExpenseRepositoryInterface
}

func NewExpenseService(repo repositories.ExpenseRepositoryInterface) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) CreateExpense(ctx context.Context, userID uint, req dto.CreateExpenseRequest) (dto.ExpenseResponse, error) {
	req_expense := models.Expense{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}
	res_expense, err := s.repo.CreateExpense(ctx, req_expense)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	return dto.ExpenseResponse{
		ID:           res_expense.ID,
		CategoryID:   res_expense.CategoryID,
		CategoryName: res_expense.CategoryName,
		Amount:       res_expense.Amount,
		Description:  &res_expense.Description,
		CreatedAt:    req_expense.CreatedAt,
	}, nil
}

func (s *ExpenseService) GetUserExpense(ctx context.Context, userID uint, expenseID int) (dto.ExpenseResponse, error) {
	res_expense, err := s.repo.GetExpenseByID(ctx, userID, uint(expenseID))
	if err != nil {
		return dto.ExpenseResponse{}, nil
	}

	return dto.ExpenseResponse{
		ID:           res_expense.ID,
		CategoryID:   res_expense.CategoryID,
		CategoryName: res_expense.CategoryName,
		Amount:       res_expense.Amount,
		Description:  &res_expense.Description,
		CreatedAt:    res_expense.CreatedAt,
	}, nil
}

func (s *ExpenseService) GetUserExpenses(ctx context.Context, userID uint) ([]dto.ExpenseResponse, error) {
	req_expenses, err := s.repo.GetExpensesByUserID(ctx, userID)
	if err != nil {
		return []dto.ExpenseResponse{}, err
	}
	res_expenses := make([]dto.ExpenseResponse, 0, len(req_expenses))
	for _, expense := range req_expenses {
		res_expenses = append(res_expenses, dto.ExpenseResponse{
			ID:           expense.ID,
			CategoryID:   expense.CategoryID,
			CategoryName: expense.CategoryName,
			Amount:       expense.Amount,
			Description:  &expense.Description,
			CreatedAt:    expense.CreatedAt,
		})
	}
	return res_expenses, nil
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, userID uint, expenseID int) error {
	return s.repo.DeleteExpense(ctx, userID, uint(expenseID))
}

func (s *ExpenseService) GetExpenseAnalytics(ctx context.Context, userID uint, period dto.ExpensePeriod) (dto.ExpenseAnalytics, error) {
	req, err := s.repo.GetExpensesByPeriod(ctx, userID, period.Period)
	if err != nil {
		return dto.ExpenseAnalytics{}, err
	}
	var total_amount float64
	total_count := len(req)
	for _, value := range req {
		total_amount += value.Amount
	}
	total_average_expense := total_amount / float64(total_count)

	largest_expense, err := s.repo.GetLargestExpenseByPeriod(ctx, userID, period.Period)
	if err != nil {
		return dto.ExpenseAnalytics{}, err
	}
	smallest_expense, err := s.repo.GetSmallestExpenseByPeriod(ctx, userID, period.Period)
	if err != nil {
		return dto.ExpenseAnalytics{}, err
	}
	var timedist float64
	if period.Period == "weekly" {
		timedist = 7
	} else if period.Period == "monthly" {
		timedist = 30
	} else if period.Period == "yearly" {
		timedist = 365
	} else {
		timedist = 0
	}
	averagePerDay := total_amount / timedist

	return dto.ExpenseAnalytics{
		Period:        period.Period,
		TotalAmount:   total_amount,
		ExpensesCount: total_count,
		AveragePerDay: averagePerDay,
		LargestExpense: dto.ExpenseResponse{
			ID:           largest_expense.ID,
			CategoryID:   largest_expense.CategoryID,
			CategoryName: largest_expense.CategoryName,
			Amount:       largest_expense.Amount,
			Description:  &largest_expense.Description,
			CreatedAt:    largest_expense.CreatedAt,
		},
		SmallestExpense: dto.ExpenseResponse{
			ID:           smallest_expense.ID,
			CategoryID:   smallest_expense.CategoryID,
			CategoryName: smallest_expense.CategoryName,
			Amount:       smallest_expense.Amount,
			Description:  &smallest_expense.Description,
			CreatedAt:    smallest_expense.CreatedAt,
		},
		AverageExpenseAmount: total_average_expense,
	}, nil
}
