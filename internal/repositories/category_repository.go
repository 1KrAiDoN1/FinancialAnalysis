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
	query := `INSERT INTO categories (user_id, name) VALUES ($1, $2) RETURNING id, name, created_at`
	result, err := c.storage.CreateCategory(ctx, query, category)
	if err != nil {
		return models.Category{}, err
	}
	return result, nil

}

func (c *CategoryRepository) GetCategoryByID(ctx context.Context, userId uint, category_id int) (models.Category, error) {
	query := `
        SELECT 
            c.id, 
            c.name, 
            c.created_at,
            COUNT(e.id) AS expense_count,
            COALESCE(SUM(e.amount), 0) AS total_amount
        FROM 
            categories c
        LEFT JOIN 
            expenses e ON c.id = e.category_id AND e.user_id = $2
        WHERE 
            c.id = $1 AND c.user_id = $2
        GROUP BY 
            c.id`
	result, err := c.storage.GetCategoryByID(ctx, query, userId, category_id)
	if err != nil {
		return models.Category{}, err
	}
	return result, nil

}

func (c *CategoryRepository) GetCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	return nil, nil
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, userID uint, category_id int) error {
	return nil
}

func (c *CategoryRepository) GetMostUsedCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	return nil, nil
}

func (c *CategoryRepository) GetTotalAmountInCategory(ctx context.Context, userID uint, categoryID int, period string) (float64, error) {
	return 0, nil
}

func (c *CategoryRepository) GetLargestExpenseInCategory(ctx context.Context, userID uint, categoryID int, period string) (models.Expense, error) {
	return models.Expense{}, nil
}

func (c *CategoryRepository) GetSmallestExpenseInCategory(ctx context.Context, userID uint, categoryID int, period string) (models.Expense, error) {
	return models.Expense{}, nil
}

func (c *CategoryRepository) GetExpenseCountInCategory(ctx context.Context, userID uint, categoryID int, period string) (int, error) {
	return 0, nil
}
