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
}

type CategoryStorageInterface interface {
	CreateCategory(ctx context.Context, query string, category models.Category) (models.Category, error)
	GetCategoryByID(ctx context.Context, query string, userID uint, categoryID int) (models.Category, error)
}

type ExpenseStorageInterface interface {
	CreateExpense(ctx context.Context, query string, expense models.Expense) (models.Expense, error)
}

type UserStorageInterface interface {
	DeleteUser(ctx context.Context, query string, userID uint) error
	GetUserStats(ctx context.Context, query string, userID uint) (models.UserStats, error)
	GetProfile(ctx context.Context, query string, userID uint) (models.User, error)
}
