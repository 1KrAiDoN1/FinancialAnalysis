package storage

import (
	"context"
	"database/sql"
	"finance/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthStorage struct {
	pool *pgxpool.Pool
}

func NewAuthStorage(pool *pgxpool.Pool) *AuthStorage {
	return &AuthStorage{
		pool: pool,
	}
}

func (s *AuthStorage) CreateUser(ctx context.Context, query string, first_name string, last_name string, email string, password string, timeOfRegistration time.Time) (models.User, error) {
	var result models.User
	err := s.pool.QueryRow(ctx, query, first_name, last_name, email, password, timeOfRegistration).Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (s *AuthStorage) CheckUserVerification(ctx context.Context, query string, email string, hashpassword string) (models.User, error) {
	var result models.User
	err := s.pool.QueryRow(ctx, query, email, hashpassword).Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (s *AuthStorage) UserExistsByEmail(ctx context.Context, query string, email string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *AuthStorage) GetUserIDbyRefreshToken(ctx context.Context, query string, refreshToken string) (int, error) {
	var userID int
	err := s.pool.QueryRow(ctx, query, refreshToken).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // токен не найден или истек
		}
		return 0, err
	}
	return userID, nil
}

func (s *AuthStorage) RemoveOldRefreshToken(ctx context.Context, query string, userID int) error {
	_, err := s.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthStorage) SaveNewRefreshToken(ctx context.Context, query string, user_id int, token models.RefreshToken) error {
	_, err := s.pool.Exec(ctx, query, user_id, token.Token, token.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}
