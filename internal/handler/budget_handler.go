package handler

import (
	"finance/internal/services"

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
	//....
}

func (b *BudgetHandler) GetBudgets(c *gin.Context) {
	//....
}

func (b *BudgetHandler) UpdateBudget(c *gin.Context) {
	//....
}

func (b *BudgetHandler) DeleteBudget(c *gin.Context) {
	//....
}
