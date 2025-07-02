package storage

import (
	"context"
	"finance/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryStorage struct {
	pool *pgxpool.Pool
}

func NewCategoryStorage(pool *pgxpool.Pool) *CategoryStorage {
	return &CategoryStorage{
		pool: pool,
	}
}

func (c *CategoryStorage) CreateCategory(ctx context.Context, query string, category models.Category) (models.Category, error) {
	var newCategory models.Category
	err := c.pool.QueryRow(ctx, query, category.UserID, category.Name).Scan(
		&newCategory.ID,
		&newCategory.Name,
		&newCategory.CreatedAt,
	)
	if err != nil {
		return models.Category{}, fmt.Errorf("failed to create category: %w", err)
	}
	return newCategory, nil
}

func (c *CategoryStorage) GetCategoryByID(ctx context.Context, query string, userID uint, categoryID int) (models.Category, error) {
	var category models.Category

	err := c.pool.QueryRow(ctx, query, categoryID, userID).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.ExpenseCount,
		&category.TotalAmount,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Category{}, fmt.Errorf("category not found")
		}
		return models.Category{}, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (c *CategoryStorage) GetCategories(ctx context.Context, query string, userID uint) ([]models.Category, error) {
	rows, err := c.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return categories, nil
}

func (c *CategoryStorage) DeleteCategory(ctx context.Context, query string, userID uint, categoryID int) error {
	result, err := c.pool.Exec(ctx, query, categoryID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("category not found or not owned by user")
	}

	return nil
}

func (c *CategoryStorage) GetMostUsedCategories(ctx context.Context, query string, userID uint) ([]models.Category, error) {
	rows, err := c.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get most used categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.CreatedAt,
			&category.ExpenseCount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategoryStorage) GetTotalAmountInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (float64, error) {
	var total float64
	err := c.pool.QueryRow(ctx, query, userID, categoryID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total amount: %w", err)
	}

	return total, nil
}

func (c *CategoryStorage) GetLargestExpenseInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error) {
	var expense models.Expense
	err := c.pool.QueryRow(ctx, query, userID, categoryID).Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.Amount,
		&expense.Description,
		&expense.Date,
		&expense.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Expense{}, nil
		}
		return models.Expense{}, fmt.Errorf("failed to get largest expense: %w", err)
	}

	return expense, nil
}

func (c *CategoryStorage) GetSmallestExpenseInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error) {
	var expense models.Expense
	err := c.pool.QueryRow(ctx, query, userID, categoryID).Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.Amount,
		&expense.Description,
		&expense.Date,
		&expense.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Expense{}, nil
		}
		return models.Expense{}, fmt.Errorf("failed to get smallest expense: %w", err)
	}

	return expense, nil
}

func (c *CategoryStorage) GetExpenseCountInCategory(ctx context.Context, query string, userID uint, categoryID int, period string) (int, error) {
	var count int
	err := c.pool.QueryRow(ctx, query, userID, categoryID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get expense count: %w", err)
	}

	return count, nil
}
