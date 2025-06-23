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

func (c *CategoryRepository) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	return models.Category{}, nil
}

func (c *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (models.Category, error) {
	return models.Category{}, nil
}

func (c *CategoryRepository) GetCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	return nil, nil
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	return nil
}

func (c *CategoryRepository) GetMostUsedCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	return nil, nil
}
