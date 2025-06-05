package services

import (
	"context"
	"finance/internal/dto"
	"finance/internal/repositories"
)

type CategoryService struct {
	repo repositories.CategoryRepositoryInterface
}

func NewCategoryService(repo repositories.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, userID uint, req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	return &dto.CategoryResponse{}, nil
}

func (c *CategoryService) GetUserCategories(ctx context.Context, userID uint) ([]*dto.CategoryResponse, error) {
	return []*dto.CategoryResponse{}, nil
}

func (c *CategoryService) UpdateCategory(ctx context.Context, userID uint, categoryID uint, req dto.UpdateCategoryRequest) error {
	return nil
}

func (c *CategoryService) DeleteCategory(ctx context.Context, userID uint, categoryID uint) error {
	return nil
}
