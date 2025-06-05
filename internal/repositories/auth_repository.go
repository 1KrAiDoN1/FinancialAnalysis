package repositories

import (
	"context"
	"finance/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository { //конструктор
	return &AuthRepository{
		pool: pool,
	}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) (int, error) {

	return 0, nil // логика для сохранения пользователя в базе данных и возвращение id пользователя
}

// func (r *AuthRepository) GetUser(username, password string) (models.User, error) {
// 	return models.User{}, nil // логика для получения пользователя из базы данных по username и password
// }
