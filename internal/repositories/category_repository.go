package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
)

type CategoryRepository struct {
	storage storage.CategoryStorageInterface
}

func NewCategoryRepository(storage storage.CategoryStorageInterface) *CategoryRepository { //конструктор
	return &CategoryRepository{
		storage: storage,
	}
}

func (c *CategoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {

}

func (c *CategoryRepository) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {

}

func (c *CategoryRepository) GetCategoryByUserID(ctx context.Context, userID uint) ([]*models.Category, error) {

}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, id uint) error {

}

func (c *CategoryRepository) GetMostUsedCategories(ctx context.Context, userID uint) ([]*models.Category, error) {

}
