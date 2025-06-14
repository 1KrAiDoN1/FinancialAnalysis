package handler

import (
	"context"
	"finance/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	budgetService services.BudgetServiceInterface
}

func NewBudgetHandler(budgetService services.BudgetServiceInterface) *BudgetHandler {
	return &BudgetHandler{
		budgetService: budgetService,
	}
}

func (b *BudgetHandler) CreateBudget(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
}

func (b *BudgetHandler) GetBudgets(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
}

func (b *BudgetHandler) UpdateBudget(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
}

func (b *BudgetHandler) DeleteBudget(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
}
