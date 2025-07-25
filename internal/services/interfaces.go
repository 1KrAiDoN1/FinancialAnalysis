package services

import (
	"context"
	"finance/internal/dto"
	"time"
)

type AuthServiceInterface interface {
	SignUp(ctx context.Context, req dto.RegisterRequest) (*dto.UserInfo, error)
	SignIn(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
	GenerateRefreshToken() (dto.RefreshTokenRequest, error)
	GenerateAccessToken(userID int) (dto.AccessTokenRequest, error)
	ValidateToken(ctx context.Context, req dto.AccessTokenRequest) (*dto.UserID, error)
	GetUserIDbyRefreshToken(ctx context.Context, refresh_token string) (int, error)
	RemoveOldRefreshToken(ctx context.Context, userID int) error
	SaveNewRefreshToken(ctx context.Context, user_id int, token dto.RefreshTokenRequest) error
}

type BudgetServiceInterface interface {
	CreateBudget(ctx context.Context, userID uint, category_id int, req dto.CreateBudgetRequest) (dto.BudgetResponse, error)
	GetUserBudgets(ctx context.Context, userID uint, category_id int) ([]dto.BudgetResponse, error)
	//UpdateBudget(ctx context.Context, userID uint, category_id int, budgetID int, req dto.UpdateBudgetRequest) error
	DeleteBudget(ctx context.Context, userID uint, category_id, budgetID int) error
	//CheckBudgetStatus(ctx context.Context, userID uint) ([]*dto.BudgetStatus, error)
}

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, userID uint, req dto.CreateCategoryRequest) (dto.CategoryResponse, error)
	GetUserCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, userID uint, categoryID int) (dto.CategoryResponse, error)
	GetMostUsedCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error)
	DeleteCategory(ctx context.Context, userID uint, categoryID int) error
	GetAnalyticsByCategory(ctx context.Context, userID uint, categoryID int, period dto.CategoryPeriod) (dto.CategoryAnalytics, error)
}

type ExpenseServiceInterface interface {
	CreateExpense(ctx context.Context, userID uint, category_id int, req dto.CreateExpenseRequest) (dto.ExpenseResponse, error)
	GetUserExpense(ctx context.Context, userID uint, category_id int, expenseID int) (dto.ExpenseResponse, error)
	GetUserExpenses(ctx context.Context, category_id int, userID uint) ([]dto.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, userID uint, category_id int, expenseID int) error
	GetExpenseAnalytics(ctx context.Context, userID uint, category_id int, period dto.ExpensePeriod) (dto.ExpenseAnalytics, error)
	updateBudgetsAfterExpense(ctx context.Context, userID uint, categoryID int, amount float64, expenseDate time.Time) error
}

type UserServiceInterface interface {
	GetProfile(ctx context.Context, userID uint) (dto.UserProfile, error)
	DeleteAccount(ctx context.Context, userID uint) error
	GetUserStats(ctx context.Context, userID uint) (dto.UserStats, error)
}
