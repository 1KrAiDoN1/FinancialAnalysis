package storage

import (
	"context"
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
