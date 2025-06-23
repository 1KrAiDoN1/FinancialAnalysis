package repositories

import (
	"context"
	"finance/internal/models"
)

type AuthRepositoryInterface interface {
	// Операции с пользователями
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
	GetUserIDbyRefreshToken(refresh_token string) (int, error)
	RemoveOldRefreshToken(userID int) error
	SaveNewRefreshToken(token models.RefreshToken) error
	CheckUserVerification(ctx context.Context, email string, hashpassword string) (*models.User, error)
	// Проверка существования
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
}

// UserRepository handles user data persistence
type UserRepositoryInterface interface {
	DeleteUser(ctx context.Context, userID uint) error
	GetUserStats(ctx context.Context, userID uint) (models.UserStats, error)
	GetProfile(ctx context.Context, userID uint) (models.User, error)
}

// CategoryRepository handles category data persistence
type CategoryRepositoryInterface interface {
	// Basic CRUD operations
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	GetCategoryByID(ctx context.Context, id int) (models.Category, error)
	GetCategories(ctx context.Context, userID uint) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id int) error
	// Additional methods
	GetMostUsedCategories(ctx context.Context, userID uint) ([]models.Category, error)
}

// ExpenseRepository handles expense data persistence
type ExpenseRepositoryInterface interface {
	// Basic CRUD operations
	CreateExpense(ctx context.Context, expense *models.Expense) error
	GetExpenseByID(ctx context.Context, id uint) (*models.Expense, error)
	GetExpenseByUserID(ctx context.Context, userID uint, expense *models.Expense) ([]*models.Expense, error)
	DeleteExpense(ctx context.Context, id uint) error
	// Analytics and reporting methods
	GetExpensesByCategory(ctx context.Context, userID uint, categoryID uint, limit, offset int) ([]*models.Expense, error)
	// Aggregation methods
	GetAverageExpenseAmount(ctx context.Context, userID uint) (float64, error)
	GetLargestExpense(ctx context.Context, userID uint) (*models.Expense, error)
	GetSmallestExpense(ctx context.Context, userID uint) (*models.Expense, error)
}

// BudgetRepository handles budget data persistence
type BudgetRepositoryInterface interface {
	// Basic CRUD operations
	CreateBudget(ctx context.Context, budget models.Budget) (models.Budget, error)
	GetBudgetByID(ctx context.Context, id int) (models.Budget, error)
	GetUserBudgets(ctx context.Context, userID uint) ([]models.Budget, error)
	GetBudgetByUserID(ctx context.Context, userID uint) ([]models.Budget, error)
	UpdateBudget(ctx context.Context, budget models.Budget) error
	DeleteBudget(ctx context.Context, userID uint, id uint) error
}
