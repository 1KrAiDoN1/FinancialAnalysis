package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/models"
	"finance/internal/repositories"
	"time"
)

type CategoryService struct {
	repo repositories.CategoryRepositoryInterface
}

func NewCategoryService(repo repositories.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, userID uint, req dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	category_req := models.Category{
		Name:      req.Name,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	category_res, err := c.repo.CreateCategory(ctx, category_req)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:        category_res.ID,
		Name:      category_res.Name,
		CreatedAt: category_res.CreatedAt,
	}, nil
}

func (c *CategoryService) GetUserCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error) {
	categories, err := c.repo.GetCategories(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		res = append(res, dto.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
		})
	}
	return res, nil
}

func (c *CategoryService) GetCategoryByID(ctx context.Context, userID uint, categoryID int) (dto.CategoryResponse, error) {
	category, err := c.repo.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}, nil
}

func (c *CategoryService) GetMostUsedCategories(ctx context.Context, userID uint) ([]dto.CategoryResponse, error) {
	categories, err := c.repo.GetMostUsedCategories(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		res = append(res, dto.CategoryResponse{
			ID:            category.ID,
			Name:          category.Name,
			CreatedAt:     category.CreatedAt,
			ExpensesCount: len(category.Expenses),
			TotalAmount:   0,
		})
	}
	return res, nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, userID uint, categoryID int) error {
	return c.repo.DeleteCategory(ctx, categoryID)
}
