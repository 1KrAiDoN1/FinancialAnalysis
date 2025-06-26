package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/models"
	"finance/internal/repositories"
	"time"
)

type CategoryService struct {
	repo repositories.CategoryRepositoryInterface
}

func NewCategoryService(repo repositories.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, userID uint, req dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	category_req := models.Category{
		Name:      req.Name,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	category_res, err := c.repo.CreateCategory(ctx, category_req)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:        category_res.ID,
		Name:      category_res.Name,
		CreatedAt: category_res.CreatedAt,
	}, nil
}

func (c *CategoryService) GetUserCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error) {
	categories, err := c.repo.GetCategories(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		res = append(res, dto.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
		})
	}
	return res, nil
}

func (c *CategoryService) GetCategoryByID(ctx context.Context, userID uint, categoryID int) (dto.CategoryResponse, error) {
	category, err := c.repo.GetCategoryByID(ctx, userID, categoryID)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}, nil
}

func (c *CategoryService) GetMostUsedCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error) {
	categories, err := c.repo.GetMostUsedCategories(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		res = append(res, dto.CategoryResponse{
			ID:            category.ID,
			Name:          category.Name,
			CreatedAt:     category.CreatedAt,
			ExpensesCount: len(category.Expenses),
			TotalAmount:   0,
		})
	}
	return res, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, userID uint, categoryID int) error {
	return c.repo.DeleteCategory(ctx, categoryID)
}

func (c *CategoryService) GetAnalyticsByCategory(ctx context.Context, userID uint, categoryID int, period dto.CategoryPeriod) (dto.CategoryAnalytics, error) {
	category_name, err := c.repo.GetCategoryByID(ctx, userID, categoryID)
	if err != nil {
		return dto.CategoryAnalytics{}, err
	}
	total_amount, err := c.repo.GetTotalAmountInCategory(ctx, userID, categoryID, period.Period)
	if err != nil {
		return dto.CategoryAnalytics{}, err
	}
	expense_count, err := c.repo.GetExpenseCountInCategory(ctx, userID, categoryID, period.Period)
	if err != nil {
		return dto.CategoryAnalytics{}, err
	}
	largest_expense, err := c.repo.GetLargestExpenseInCategory(ctx, userID, categoryID, period.Period)
	if err != nil {
		return dto.CategoryAnalytics{}, err
	}
	smallest_expense, err := c.repo.GetSmallestExpenseInCategory(ctx, userID, categoryID, period.Period)
	if err != nil {
		return dto.CategoryAnalytics{}, err
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

	return dto.CategoryAnalytics{
		CategoryID:    uint(categoryID),
		CategoryName:  category_name.Name,
		Period:        period.Period,
		TotalAmount:   total_amount,
		ExpensesCount: expense_count,
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
		AverageExpenseAmount: float64(expense_count) / timedist,
	}, nil
}
