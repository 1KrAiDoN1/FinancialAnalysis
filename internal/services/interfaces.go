package services

import (
	"context"
	"finance/internal/dto"
)

type AuthServiceInterface interface {
	SignUp(ctx context.Context, req dto.RegisterRequest) (*dto.UserInfo, error)
	SignIn(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
	Logout(ctx context.Context, req dto.LogoutRequest) error
	GenerateRefreshToken() (*dto.AuthResponse, error)
	GenerateAccessToken(req dto.LoginRequest) (*dto.AuthResponse, error)
	ValidateToken(ctx context.Context, req dto.AccessTokenRequest) (*dto.UserInfo, error)
	// ValidateToken(ctx context.Context, token string) (*models.User, error)
}

type BudgetServiceInterface interface {
	CreateBudget(ctx context.Context, userID uint, req dto.CreateBudgetRequest) (dto.BudgetResponse, error)
	GetUserBudgets(ctx context.Context, userID uint) ([]dto.BudgetResponse, error)
	UpdateBudget(ctx context.Context, userID uint, budgetID int, req dto.UpdateBudgetRequest) error
	DeleteBudget(ctx context.Context, userID uint, budgetID int) error
	//CheckBudgetStatus(ctx context.Context, userID uint) ([]*dto.BudgetStatus, error)
}

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, userID uint, req dto.CreateCategoryRequest) (dto.CategoryResponse, error)
	GetUserCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, userID uint, categoryID int) (dto.CategoryResponse, error)
	DeleteCategory(ctx context.Context, userID uint, categoryID int) error
}

type ExpenseServiceInterface interface {
	CreateExpense(ctx context.Context, userID uint, req dto.CreateExpenseRequest) (dto.ExpenseResponse, error)
	GetUserExpense(ctx context.Context, userID uint, expenseID int) (dto.ExpenseResponse, error)
	GetUserExpenses(ctx context.Context, userID uint) ([]dto.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, userID uint, expenseID int) error
	GetExpenseAnalytics(ctx context.Context, userID uint, period dto.ExpensePeriod) (dto.ExpenseAnalytics, error)
}

type UserServiceInterface interface {
	GetProfile(ctx context.Context, userID uint) (dto.UserProfile, error)
	DeleteAccount(ctx context.Context, userID uint) error
	GetUserStats(ctx context.Context, userID uint) (dto.UserStats, error)
}
