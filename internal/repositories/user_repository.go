package repositories

import (
	"context"

	"finance/internal/models"
	storage "finance/internal/storages"
)

type UserRepository struct {
	storage storage.UserStorageInterface
}

func NewUserRepository(storage storage.UserStorageInterface) *UserRepository { //конструктор
	return &UserRepository{
		storage: storage,
	}
}

func (u *UserRepository) DeleteUser(ctx context.Context, userID uint) error {
	query := `DELETE FROM users WHERE id = $1`
	err := u.storage.DeleteUser(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserStats(ctx context.Context, userID uint) (models.UserStats, error) {
	query := `SELECT COUNT(DISTINCT e.id) as total_expenses_count, 
	COUNT(DISTINCT c.id) as total_categories_count, 
	COUNT(DISTINCT b.id) as total_budgets_count, 
	COALESCE(SUM(CASE WHEN e.date >= (CURRENT_TIMESTAMP - INTERVAL '30 days') 
	THEN e.amount 
	ELSE 0 
    END), 0) as monthly_expenses_sum,
    COALESCE(SUM(CASE 
        WHEN e.date >= (CURRENT_TIMESTAMP - INTERVAL '7 days')
        THEN e.amount 
        ELSE 0 
    END), 0) as weekly_expenses_sum

	FROM users u
	LEFT JOIN categories c ON u.id = c.user_id
	LEFT JOIN expenses e ON u.id = e.user_id
	LEFT JOIN budgets b ON u.id = b.user_id
	WHERE u.id = $1
	GROUP BY u.id;`
	result, err := u.storage.GetUserStats(ctx, query, userID)
	if err != nil {
		return models.UserStats{}, err
	}
	return result, nil
}

func (u *UserRepository) GetProfile(ctx context.Context, userID uint) (models.User, error) {
	query := `SELECT id, first_name, last_name, email, time_of_registration FROM users WHERE id = $1`
	result, err := u.storage.GetProfile(ctx, query, userID)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}
