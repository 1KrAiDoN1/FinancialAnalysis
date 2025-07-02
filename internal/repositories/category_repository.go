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
	query := `SELECT id, user_id, name, created_at FROM categories WHERE user_id = $1 ORDER BY created_at DESC`
	result, err := c.storage.GetCategories(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, userID uint, category_id int) error {
	query := `DELETE FROM categories WHERE id = $1 AND user_id = $2`
	err := c.storage.DeleteCategory(ctx, query, userID, category_id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryRepository) GetMostUsedCategories(ctx context.Context, userID uint) ([]models.Category, error) {
	query := `
		SELECT c.id, c.user_id, c.name, c.created_at, COUNT(e.id) as expense_count
		FROM categories c
		LEFT JOIN expenses e ON c.id = e.category_id AND e.user_id = $1
		WHERE c.user_id = $1
		GROUP BY c.id
		ORDER BY expense_count DESC
		LIMIT 5`
	result, err := c.storage.GetMostUsedCategories(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CategoryRepository) GetTotalAmountInCategory(ctx context.Context, userID uint, categoryID int, period string) (float64, error) {
	query := `SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id = $1 AND category_id = $2`

	switch period {
	case "weekly":
		query += " AND date >= date_trunc('week', CURRENT_DATE)"
	case "monthly":
		query += " AND date >= date_trunc('month', CURRENT_DATE)"
	case "yearly":
		query += " AND date >= date_trunc('year', CURRENT_DATE)"
	}
	result, err := c.storage.GetTotalAmountInCategory(ctx, query, userID, categoryID, period)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *CategoryRepository) GetLargestExpenseInCategory(ctx context.Context, userID uint, categoryID int, period string) (models.Expense, error) {
	query := `SELECT id, user_id, category_id, amount, description, date, created_at FROM expenses WHERE user_id = $1 AND category_id = $2`

	switch period {
	case "weekly":
		query += " AND date >= date_trunc('week', CURRENT_DATE)"
	case "monthly":
		query += " AND date >= date_trunc('month', CURRENT_DATE)"
	case "yearly":
		query += " AND date >= date_trunc('year', CURRENT_DATE)"
	}

	query += " ORDER BY amount DESC LIMIT 1"
	result, err := c.storage.GetLargestExpenseInCategory(ctx, query, userID, categoryID, period)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil

}

func (c *CategoryRepository) GetSmallestExpenseInCategory(ctx context.Context, userID uint, categoryID int, period string) (models.Expense, error) {
	query := `SELECT id, user_id, category_id, amount, description, date, created_at FROM expenses WHERE user_id = $1 AND category_id = $2 AND amount > 0`

	switch period {
	case "weekly":
		query += " AND date >= date_trunc('week', CURRENT_DATE)"
	case "monthly":
		query += " AND date >= date_trunc('month', CURRENT_DATE)"
	case "yearly":
		query += " AND date >= date_trunc('year', CURRENT_DATE)"
	}

	query += " ORDER BY amount ASC LIMIT 1"
	result, err := c.storage.GetSmallestExpenseInCategory(ctx, query, userID, categoryID, period)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil
}

func (c *CategoryRepository) GetExpenseCountInCategory(ctx context.Context, userID uint, categoryID int, period string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM expenses
		WHERE user_id = $1 AND category_id = $2`

	switch period {
	case "weekly":
		query += " AND date >= date_trunc('week', CURRENT_DATE)"
	case "monthly":
		query += " AND date >= date_trunc('month', CURRENT_DATE)"
	case "yearly":
		query += " AND date >= date_trunc('year', CURRENT_DATE)"
	}
	result, err := c.storage.GetExpenseCountInCategory(ctx, query, userID, categoryID, period)
	if err != nil {
		return 0, err
	}
	return result, nil
}
