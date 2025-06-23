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
	return nil
}

func (u *UserRepository) GetUserStats(ctx context.Context, userID uint) (models.UserStats, error) {
	return models.UserStats{}, nil
}

func (u *UserRepository) GetProfile(ctx context.Context, userID uint) (models.User, error) {
	return models.User{}, nil
}
