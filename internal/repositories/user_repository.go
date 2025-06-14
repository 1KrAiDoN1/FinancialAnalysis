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

func (u *UserRepository) DeleteUser(ctx context.Context, id uint) error {

}

func (u *UserRepository) GetUserStats(ctx context.Context, userID uint) (*models.UserStats, error) {

}
