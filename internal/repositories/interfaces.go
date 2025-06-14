package repositories

import (
	"context"
	"finance/internal/models"
	"time"
)

type AuthRepositoryInterface interface {
	// Операции с пользователями
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
	// Проверка существования
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
}

// UserRepository handles user data persistence
type UserRepositoryInterface interface {
	DeleteUser(ctx context.Context, id uint) error
	GetUserStats(ctx context.Context, userID uint) (*models.UserStats, error)
}

// CategoryRepository handles category data persistence
type CategoryRepositoryInterface interface {
	// Basic CRUD operations
	CreateCategory(ctx context.Context, category *models.Category) error
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryByUserID(ctx context.Context, userID uint) ([]*models.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
	// Additional methods
	GetMostUsedCategories(ctx context.Context, userID uint) ([]*models.Category, error)
}

// ExpenseRepository handles expense data persistence
type ExpenseRepositoryInterface interface {
	// Basic CRUD operations
	CreateExpense(ctx context.Context, expense *models.Expense) error
	GetExpenseByID(ctx context.Context, id uint) (*models.Expense, error)
	GetExpenseByUserID(ctx context.Context, userID uint, expense *models.Expense) ([]*models.Expense, error)
	DeleteExpense(ctx context.Context, id uint) error
	// Analytics and reporting methods
	GetExpensesByDateRange(ctx context.Context, userID uint, startDate, endDate time.Time) ([]*models.Expense, error)
	GetExpensesByCategory(ctx context.Context, userID uint, categoryID uint, limit, offset int) ([]*models.Expense, error)
	// Aggregation methods
	GetAverageExpenseAmount(ctx context.Context, userID uint) (float64, error)
	GetLargestExpense(ctx context.Context, userID uint) (*models.Expense, error)
	GetSmallestExpense(ctx context.Context, userID uint) (*models.Expense, error)
}

// BudgetRepository handles budget data persistence
type BudgetRepositoryInterface interface {
	// Basic CRUD operations
	CreateBudget(ctx context.Context, budget *models.Budget) error
	GetBudgetByID(ctx context.Context, id uint) (*models.Budget, error)
	GetBudgetByUserID(ctx context.Context, userID uint) ([]*models.Budget, error)
	DeleteBudget(ctx context.Context, id uint) error
}
