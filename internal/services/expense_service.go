package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/models"
	repositories "finance/internal/repositories"
	"time"
)

type ExpenseService struct {
	repo        repositories.ExpenseRepositoryInterface
	budget_repo repositories.BudgetRepositoryInterface
}

func NewExpenseService(repo repositories.ExpenseRepositoryInterface, budget_repo repositories.BudgetRepositoryInterface) *ExpenseService {
	return &ExpenseService{
		repo:        repo,
		budget_repo: budget_repo,
	}
}

func (s *ExpenseService) CreateExpense(ctx context.Context, userID uint, req dto.CreateExpenseRequest) (dto.ExpenseResponse, error) {
	req_expense := models.Expense{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        req.Date,
		CreatedAt:   time.Now(),
	}
	res_expense, err := s.repo.CreateExpense(ctx, req_expense)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	// Обновляем бюджеты после создания расхода
	err = s.updateBudgetsAfterExpense(ctx, userID, int(req.CategoryID), req.Amount, req.Date)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	return dto.ExpenseResponse{
		ID:           res_expense.ID,
		CategoryID:   res_expense.CategoryID,
		CategoryName: res_expense.CategoryName,
		Amount:       res_expense.Amount,
		Description:  &res_expense.Description,
		Date:         res_expense.Date,
		CreatedAt:    req_expense.CreatedAt,
	}, nil
}

func (s *ExpenseService) GetUserExpense(ctx context.Context, userID uint, category_id int, expenseID int) (dto.ExpenseResponse, error) {
	res_expense, err := s.repo.GetExpenseByID(ctx, userID, category_id, uint(expenseID))
	if err != nil {
		return dto.ExpenseResponse{}, nil
	}

	return dto.ExpenseResponse{
		ID:           res_expense.ID,
		CategoryID:   res_expense.CategoryID,
		CategoryName: res_expense.CategoryName,
		Amount:       res_expense.Amount,
		Description:  &res_expense.Description,
		Date:         res_expense.Date,
		CreatedAt:    res_expense.CreatedAt,
	}, nil
}

func (s *ExpenseService) GetUserExpenses(ctx context.Context, category_id int, userID uint) ([]dto.ExpenseResponse, error) {
	req_expenses, err := s.repo.GetExpensesByUserID(ctx, category_id, userID)
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
			Date:         expense.Date,
			CreatedAt:    expense.CreatedAt,
		})
	}
	return res_expenses, nil
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, userID uint, category_id int, expenseID int) error {
	// Получаем информацию о расходе перед удалением для возврата бюджета
	expense, err := s.repo.GetExpenseByID(ctx, userID, category_id, uint(expenseID))
	if err != nil {
		return err
	}

	err = s.repo.DeleteExpense(ctx, userID, category_id, uint(expenseID))
	if err != nil {
		return err
	}

	// Возвращаем средства в бюджеты после удаления расхода
	err = s.restoreBudgetsAfterExpenseDeletion(ctx, userID, int(expense.CategoryID), expense.Amount, expense.Date)
	if err != nil {
		// Логируем ошибку, но не прерываем процесс удаления расхода
	}

	return nil

}

func (s *ExpenseService) GetExpenseAnalytics(ctx context.Context, userID uint, category_id int, period dto.ExpensePeriod) (dto.ExpenseAnalytics, error) {
	req, err := s.repo.GetExpensesByPeriod(ctx, userID, category_id, period.Period)
	if err != nil {
		return dto.ExpenseAnalytics{}, err
	}
	var total_amount float64
	total_count := len(req)
	for _, value := range req {
		total_amount += value.Amount
	}
	total_average_expense := total_amount / float64(total_count)

	largest_expense, err := s.repo.GetLargestExpenseByPeriod(ctx, userID, category_id, period.Period)
	if err != nil {
		return dto.ExpenseAnalytics{}, err
	}
	smallest_expense, err := s.repo.GetSmallestExpenseByPeriod(ctx, userID, category_id, period.Period)
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
			Date:         largest_expense.Date,
			CreatedAt:    largest_expense.CreatedAt,
		},
		SmallestExpense: dto.ExpenseResponse{
			ID:           smallest_expense.ID,
			CategoryID:   smallest_expense.CategoryID,
			CategoryName: smallest_expense.CategoryName,
			Amount:       smallest_expense.Amount,
			Description:  &smallest_expense.Description,
			Date:         smallest_expense.Date,
			CreatedAt:    smallest_expense.CreatedAt,
		},
		AverageExpenseAmount: total_average_expense,
	}, nil
}

// updateBudgetsAfterExpense обновляет все подходящие бюджеты после добавления расхода
func (s *ExpenseService) updateBudgetsAfterExpense(ctx context.Context, userID uint, categoryID int, amount float64, expenseDate time.Time) error {
	// Получаем все активные бюджеты пользователя для данной категории
	budgets, err := s.budget_repo.GetActiveBudgetsByCategoryAndDate(ctx, userID, categoryID, expenseDate)
	if err != nil {
		return err
	}

	// Обновляем каждый подходящий бюджет
	for _, budget := range budgets {
		err = s.budget_repo.UpdateSpentAmount(ctx, categoryID, budget.ID, budget.SpentAmount+amount)
		if err != nil {
			return err
		}
	}

	return nil
}

// restoreBudgetsAfterExpenseDeletion возвращает средства в бюджеты после удаления расхода
func (s *ExpenseService) restoreBudgetsAfterExpenseDeletion(ctx context.Context, userID uint, categoryID int, amount float64, expenseDate time.Time) error {
	// Получаем все активные бюджеты пользователя для данной категории
	budgets, err := s.budget_repo.GetActiveBudgetsByCategoryAndDate(ctx, userID, categoryID, expenseDate)
	if err != nil {
		return err
	}

	// Возвращаем средства в каждый подходящий бюджет
	for _, budget := range budgets {
		newSpentAmount := budget.SpentAmount - amount
		if newSpentAmount < 0 {
			newSpentAmount = 0
		}
		err = s.budget_repo.UpdateSpentAmount(ctx, categoryID, budget.ID, newSpentAmount)
		if err != nil {
			return err
		}
	}

	return nil
}
