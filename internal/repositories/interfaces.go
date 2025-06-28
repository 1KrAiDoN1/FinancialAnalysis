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
	GetUserIDbyRefreshToken(refresh_token string) (int, error)
	RemoveOldRefreshToken(userID int) error
	SaveNewRefreshToken(token models.RefreshToken) error
	CheckUserVerification(ctx context.Context, email string, hash_password string) (*models.User, error)
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
	GetCategoryByID(ctx context.Context, userId uint, category_id int) (models.Category, error)
	GetCategories(ctx context.Context, userID uint) ([]models.Category, error)
	DeleteCategory(ctx context.Context, userID uint, category_id int) error
	// Additional methods
	GetMostUsedCategories(ctx context.Context, userID uint) ([]models.Category, error)
	GetTotalAmountInCategory(ctx context.Context, userID uint, categoryID int, period string) (float64, error)
	GetLargestExpenseInCategory(ctx context.Context, userID uint, categoryID int, period string) (models.Expense, error)
	GetSmallestExpenseInCategory(ctx context.Context, userID uint, categoryID int, period string) (models.Expense, error)
	GetExpenseCountInCategory(ctx context.Context, userID uint, categoryID int, period string) (int, error)
}

// ExpenseRepository handles expense data persistence
type ExpenseRepositoryInterface interface {
	// Basic CRUD operations
	CreateExpense(ctx context.Context, expense models.Expense) (models.Expense, error)
	GetExpenseByID(ctx context.Context, userID uint, category_id int, expense_id uint) (models.Expense, error)
	GetExpensesByUserID(ctx context.Context, category_id int, userID uint) ([]models.Expense, error)
	GetExpensesByPeriod(ctx context.Context, userID uint, category_id int, period string) ([]models.Expense, error)
	DeleteExpense(ctx context.Context, userID uint, category_id int, id uint) error
	DeleteExpensesInCategory(ctx context.Context, userID uint, categoryID int) error
	// Analytics and reporting methods
	GetExpensesByCategory(ctx context.Context, userID uint, categoryID int, limit, offset int) ([]models.Expense, error)
	GetExpensesByCategoryAndPeriod(ctx context.Context, userID uint, categoryID int, startDate, endDate time.Time) ([]models.Expense, error)
	// Aggregation methods
	GetLargestExpenseByPeriod(ctx context.Context, userID uint, category_id int, period string) (models.Expense, error)
	GetSmallestExpenseByPeriod(ctx context.Context, userID uint, category_id int, period string) (models.Expense, error)
}

// BudgetRepository handles budget data persistence
type BudgetRepositoryInterface interface {
	// Basic CRUD operations
	CreateBudget(ctx context.Context, budget models.Budget) (models.Budget, error)
	GetBudgetByID(ctx context.Context, userID uint, category_id int, budget_id int) (models.Budget, error)
	GetUserBudgets(ctx context.Context, category_id int, userID uint) ([]models.Budget, error)
	//GetBudgetByUserID(ctx context.Context, category_id int, userID uint) ([]models.Budget, error)
	//UpdateBudget(ctx context.Context, budget models.Budget) error
	DeleteBudget(ctx context.Context, userID uint, category_id int, budget_id int) error
	DeleteBudgetsInCategory(ctx context.Context, userID uint, categoryID int) error
	UpdateSpentAmount(ctx context.Context, category_id int, budgetID uint, spentAmount float64) error
	GetActiveBudgetsByCategoryAndDate(ctx context.Context, userID uint, categoryID int, date time.Time) ([]models.Budget, error)
}
