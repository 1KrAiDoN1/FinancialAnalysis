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

// func (c *CategoryStorage)

// func (c *CategoryStorage)

// func (c *CategoryStorage)

// func (c *CategoryStorage)

// func (c *CategoryStorage)

// func (c *CategoryStorage)

// func (c *CategoryStorage)
