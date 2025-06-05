package repositories

import (
	"context"
	"finance/internal/dto"
	"finance/internal/models"
	"time"
)

type AuthRepositoryInterface interface {
	// Операции с пользователями
	CreateUser(ctx context.Context, user *models.User) (int, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUserByID(ctx context.Context, userID uint) (*models.User, error)
	// UpdateUser(ctx context.Context, user *models.User) error

	// // Операции с токенами
	// SaveRefreshToken(ctx context.Context, token *models.RefreshToken) error
	// GetRefreshTokenByToken(ctx context.Context, tokenString string) (*models.RefreshToken, error)
	// GetRefreshTokenByUserID(ctx context.Context, userID uint) (*models.RefreshToken, error)
	// UpdateRefreshToken(ctx context.Context, token *models.RefreshToken) error
	// DeleteRefreshToken(ctx context.Context, tokenString string) error
	// DeleteRefreshTokensByUserID(ctx context.Context, userID uint) error

	// // Операции с access токенами (если храните в БД)
	// SaveAccessToken(ctx context.Context, token *models.AccessToken) error
	// GetAccessTokenByToken(ctx context.Context, tokenString string) (*models.AccessToken, error)
	// DeleteAccessToken(ctx context.Context, tokenString string) error
	// DeleteExpiredAccessTokens(ctx context.Context) error

	// // Проверка существования
	// UserExistsByEmail(ctx context.Context, email string) (bool, error)

	// // Операции для безопасности
	// InvalidateAllUserTokens(ctx context.Context, userID uint) error
}

// UserRepository handles user data persistence
type UserRepositoryInterface interface {
	// Basic CRUD operations
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error

	// Additional methods
	UpdateLastLogin(ctx context.Context, userID uint) error
	UpdatePassword(ctx context.Context, userID uint, hashedPassword string) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	GetUserStats(ctx context.Context, userID uint) (*dto.UserStats, error)
	List(ctx context.Context, limit, offset int) ([]*models.User, error)
	CountUsers(ctx context.Context) (int64, error)
}

// CategoryRepository handles category data persistence
type CategoryRepositoryInterface interface {
	// Basic CRUD operations
	CreateCategory(ctx context.Context, category *models.Category) error
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryByUserID(ctx context.Context, userID uint) ([]*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, id uint) error

	// Additional methods
	GetByName(ctx context.Context, userID uint, name string) (*models.Category, error)
	GetWithExpenseCount(ctx context.Context, userID uint) ([]*dto.CategoryWithStats, error)
	CheckCategoryExists(ctx context.Context, userID uint, name string) (bool, error)
	GetCategoryUsage(ctx context.Context, categoryID uint, period time.Time) (*dto.CategoryUsage, error)
	GetMostUsedCategories(ctx context.Context, userID uint, limit int) ([]*models.Category, error)
	BulkDelete(ctx context.Context, userID uint, categoryIDs []uint) error
}

// ExpenseRepository handles expense data persistence
type ExpenseRepositoryInterface interface {
	// Basic CRUD operations
	CreateExpense(ctx context.Context, expense *models.Expense) error
	GetExpenseByID(ctx context.Context, id uint) (*models.Expense, error)
	GetExpenseByUserID(ctx context.Context, userID uint, filter dto.ExpenseFilter) ([]*models.Expense, error)
	UpdateExpense(ctx context.Context, expense *models.Expense) error
	DeleteExpense(ctx context.Context, id uint) error

	// Analytics and reporting methods
	GetExpensesByDateRange(ctx context.Context, userID uint, startDate, endDate time.Time) ([]*models.Expense, error)
	GetExpensesByCategory(ctx context.Context, userID uint, categoryID uint, limit, offset int) ([]*models.Expense, error)
	GetTotalByCategory(ctx context.Context, userID uint, startDate, endDate time.Time) ([]*dto.CategoryTotal, error)
	GetTotalByPeriod(ctx context.Context, userID uint, period string) ([]*dto.PeriodTotal, error)
	GetMonthlyTotals(ctx context.Context, userID uint, year int) ([]*dto.MonthlyTotal, error)
	GetDailyTotals(ctx context.Context, userID uint, startDate, endDate time.Time) ([]*dto.DailyTotal, error)

	// Aggregation methods
	GetTotalExpenses(ctx context.Context, userID uint) (float64, error)
	GetTotalExpensesByPeriod(ctx context.Context, userID uint, startDate, endDate time.Time) (float64, error)
	GetAverageExpenseAmount(ctx context.Context, userID uint) (float64, error)
	GetLargestExpense(ctx context.Context, userID uint) (*models.Expense, error)
	GetSmallestExpense(ctx context.Context, userID uint) (*models.Expense, error)
	CountExpenses(ctx context.Context, userID uint) (int64, error)

	// Search and filter methods
	SearchExpenses(ctx context.Context, userID uint, query string, limit, offset int) ([]*models.Expense, error)
	GetRecentExpenses(ctx context.Context, userID uint, limit int) ([]*models.Expense, error)
	GetExpensesAboveAmount(ctx context.Context, userID uint, amount float64) ([]*models.Expense, error)
	GetExpensesBelowAmount(ctx context.Context, userID uint, amount float64) ([]*models.Expense, error)

	// Bulk operations
	BulkCreate(ctx context.Context, expenses []*models.Expense) error
	BulkDelete(ctx context.Context, userID uint, expenseIDs []uint) error
	DeleteByCategory(ctx context.Context, categoryID uint) error
	DeleteByDateRange(ctx context.Context, userID uint, startDate, endDate time.Time) error
}

// BudgetRepository handles budget data persistence
type BudgetRepositoryInterface interface {
	// Basic CRUD operations
	CreateBudget(ctx context.Context, budget *models.Budget) error
	GetBudgetByID(ctx context.Context, id uint) (*models.Budget, error)
	GetBudgetByUserID(ctx context.Context, userID uint) ([]*models.Budget, error)
	UpdateBudget(ctx context.Context, budget *models.Budget) error
	DeleteBudget(ctx context.Context, id uint) error

	// Budget-specific methods
	GetByCategory(ctx context.Context, userID uint, categoryID uint) ([]*models.Budget, error)
	GetActiveBudgets(ctx context.Context, userID uint) ([]*models.Budget, error)
	GetExpiredBudgets(ctx context.Context, userID uint) ([]*models.Budget, error)
	GetBudgetsByPeriod(ctx context.Context, userID uint, startDate, endDate time.Time) ([]*models.Budget, error)

	// Budget monitoring methods
	GetBudgetUsage(ctx context.Context, budgetID uint) (*dto.BudgetUsage, error)
	GetBudgetsWithUsage(ctx context.Context, userID uint) ([]*dto.BudgetWithUsage, error)
	CheckBudgetLimit(ctx context.Context, budgetID uint, amount float64) (bool, error)
	GetOverBudgetAlerts(ctx context.Context, userID uint) ([]*dto.BudgetAlert, error)
	GetNearLimitBudgets(ctx context.Context, userID uint, threshold float64) ([]*models.Budget, error)

	// Bulk operations
	BulkUpdateSpent(ctx context.Context, updates []dto.BudgetSpentUpdate) error
	DeleteByCategory(ctx context.Context, categoryID uint) error
	ArchiveExpiredBudgets(ctx context.Context) error
}

// TokenRepository handles refresh token persistence
// type TokenRepository interface {
// 	// Token CRUD operations
// 	Create(ctx context.Context, token *models.RefreshToken) error
// 	GetByToken(ctx context.Context, token string) (*models.RefreshToken, error)
// 	GetByUserID(ctx context.Context, userID uint) ([]*models.RefreshToken, error)
// 	Update(ctx context.Context, token *models.RefreshToken) error
// 	Delete(ctx context.Context, id uint) error
// 	DeleteByToken(ctx context.Context, token string) error

// 	// Token management methods
// 	DeleteByUserID(ctx context.Context, userID uint) error
// 	DeleteExpiredTokens(ctx context.Context) error
// 	IsTokenValid(ctx context.Context, token string) (bool, error)
// 	RevokeAllUserTokens(ctx context.Context, userID uint) error
// 	CleanupExpiredTokens(ctx context.Context, batchSize int) (int64, error)
// 	CountActiveTokens(ctx context.Context, userID uint) (int64, error)
// }
