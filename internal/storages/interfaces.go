package storage

import (
	"context"
	"finance/internal/models"
	"time"
)

type AuthStorageInterface interface {
	CreateUser(ctx context.Context, query string, first_name string, last_name string, email string, password string, timeOfRegistration time.Time) (models.User, error)
	CheckUserVerification(ctx context.Context, query string, email string, hashpassword string) (models.User, error)
	UserExistsByEmail(ctx context.Context, query string, email string) (bool, error)
	GetUserIDbyRefreshToken(ctx context.Context, query string, refreshToken string) (int, error)
	RemoveOldRefreshToken(ctx context.Context, query string, userID int) error
	SaveNewRefreshToken(ctx context.Context, query string, user_id int, token models.RefreshToken) error
}

type BudgetStorageInterface interface {
	CreateBudget(ctx context.Context, query string, budget models.Budget) (models.Budget, error)
	GetBudgetByID(ctx context.Context, query string, userID uint, category_id int, budget_id int) (models.Budget, error)
	GetUserBudgets(ctx context.Context, query string, category_id int, userID uint) ([]models.Budget, error)
	DeleteBudget(ctx context.Context, query string, userID uint, category_id int, budget_id int) error
	DeleteBudgetsInCategory(ctx context.Context, query string, userID uint, categoryID int) error
	UpdateSpentAmount(ctx context.Context, query string, category_id int, budgetID uint, spentAmount float64) error
	GetActiveBudgetsByCategoryAndDate(ctx context.Context, query string, userID uint, categoryID int, date time.Time) ([]models.Budget, error)
}

type CategoryStorageInterface interface {
	CreateCategory(ctx context.Context, query string, category models.Category) (models.Category, error)
	GetCategoryByID(ctx context.Context, query string, userID uint, categoryID int) (models.Category, error)
	GetCategories(ctx context.Context, query string, userID uint) ([]models.Category, error)
	DeleteCategory(ctx context.Context, query string, userID uint, categoryID int) error
	GetMostUsedCategories(ctx context.Context, query string, userID uint) ([]models.Category, error)
	GetTotalAmountInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (float64, error)
	GetLargestExpenseInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error)
	GetSmallestExpenseInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error)
	GetExpenseCountInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (int, error)
}

type ExpenseStorageInterface interface {
	CreateExpense(ctx context.Context, query string, expense models.Expense) (models.Expense, error)
	GetExpenseByID(ctx context.Context, query string, userID uint, categoryID int, id uint) (models.Expense, error)
	GetExpensesByUserID(ctx context.Context, query string, categoryID int, userID uint) ([]models.Expense, error)
	GetExpensesByPeriod(ctx context.Context, query string, userID uint, categoryID int, period string) ([]models.Expense, error)
	DeleteExpense(ctx context.Context, query string, userID uint, categoryID int, id uint) error
	DeleteExpensesInCategory(ctx context.Context, query string, userID uint, categoryID int) error
	GetExpensesByCategory(ctx context.Context, query string, userID uint, categoryID int) ([]models.Expense, error)
	GetLargestExpenseByPeriod(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error)
	GetSmallestExpenseByPeriod(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error)
	GetExpensesByCategoryAndPeriod(ctx context.Context, query string, userID uint, categoryID int, startDate, endDate time.Time) ([]models.Expense, error)
}

type UserStorageInterface interface {
	DeleteUser(ctx context.Context, query string, userID uint) error
	GetUserStats(ctx context.Context, query string, userID uint) (models.UserStats, error)
	GetProfile(ctx context.Context, query string, userID uint) (models.User, error)
}
