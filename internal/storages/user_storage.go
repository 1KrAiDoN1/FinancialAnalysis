package storage

import (
	"context"
	"errors"
	"finance/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	pool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	return &UserStorage{
		pool: pool,
	}
}

func (s *UserStorage) DeleteUser(ctx context.Context, query string, userID uint) error {
	_, err := s.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) GetUserStats(ctx context.Context, query string, userID uint) (models.UserStats, error) {
	var stats models.UserStats
	err := s.pool.QueryRow(ctx, query, userID).Scan(
		&stats.TotalExpenses,
		&stats.TotalCategories,
		&stats.TotalBudgets,
		&stats.MonthlyExpenses,
		&stats.WeeklyExpenses,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.UserStats{}, fmt.Errorf("user with id %d not found", userID)
		}
		return models.UserStats{}, fmt.Errorf("failed to get user stats: %w", err)
	}
	return stats, nil
}

func (s *UserStorage) GetProfile(ctx context.Context, query string, userID uint) (models.User, error) {
	var user_profile models.User
	err := s.pool.QueryRow(ctx, query, userID).Scan(
		&user_profile.ID,
		&user_profile.FirstName,
		&user_profile.LastName,
		&user_profile.Email,
		&user_profile.TimeOfRegistration,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, fmt.Errorf("user with id %d not found", userID)
		}
		return models.User{}, fmt.Errorf("failed to get user profile: %w", err)
	}
	return user_profile, nil

}
