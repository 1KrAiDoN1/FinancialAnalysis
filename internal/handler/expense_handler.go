package handler

import (
	"finance/internal/services"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	expenseService services.ExpenseServiceInterface
}

func NewExpenseHandler(expenseService services.ExpenseServiceInterface) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	// ....
}

func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	// ....
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	// ....
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	// ....
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	// ....
}

func (h *ExpenseHandler) GetAnalytics(c *gin.Context) {
	// ....
}
