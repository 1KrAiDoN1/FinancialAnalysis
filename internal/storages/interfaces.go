package storage

import (
	"context"
	"finance/internal/models"
	"time"
)

type AuthStorageInterface interface {
	CreateUser(ctx context.Context, query string, first_name string, last_name string, email string, password string, timeOfRegistration time.Time) (models.User, error)
	// CheckUserVerification(ctx context.Context, email string, hashpassword string) (models.User, error)
	// GetUserByID(ctx context.Context, userID uint) (models.User, error)
	// UserExistsByEmail(ctx context.Context, email string) (bool, error)
	// GetUserIDbyRefreshToken(refreshToken string) (int, error)
	// RemoveOldRefreshToken(userID int) error
	// SaveNewRefreshToken(token models.RefreshToken) error
}

type BudgetStorageInterface interface {
}

type CategoryStorageInterface interface {
}

type ExpenseStorageInterface interface {
}

type UserStorageInterface interface {
}
