package handler

import (
	"finance/internal/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryHandler(categoryService services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	// ...
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	// ...
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	// ...
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	// ...
}
