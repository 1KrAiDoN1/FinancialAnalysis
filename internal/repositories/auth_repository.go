package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
)

type AuthRepository struct {
	storage storage.AuthStorageInterface
}

func NewAuthRepository(storage storage.AuthStorageInterface) *AuthRepository { //конструктор
	return &AuthRepository{
		storage: storage,
	}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	return &models.User{}, nil // логика для сохранения пользователя в базе данных и возвращение id пользователя
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{}, nil // логика для получения пользователя по email
}

func (r *AuthRepository) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	return &models.User{}, nil // логика для получения пользователя по id
}

func (r *AuthRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	return true, nil
}
