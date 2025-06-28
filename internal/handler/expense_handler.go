package handler

import (
	"context"
	"finance/internal/dto"
	"finance/internal/middleware"
	"finance/internal/services"
	"net/http"
	"strconv"
	"time"

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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	var newexpense dto.CreateExpenseRequest
	if err := c.BindJSON(&newexpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	createdExpense, err := h.expenseService.CreateExpense(ctx, userID, newexpense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ExpenseResponse{
		ID:           createdExpense.ID,
		CategoryID:   uint(category_id),
		CategoryName: createdExpense.Category.Name,
		Amount:       createdExpense.Amount,
		Description:  createdExpense.Description,
		Date:         createdExpense.Date,
		CreatedAt:    createdExpense.CreatedAt,
	})

}

func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	expenseID, err := strconv.Atoi(c.Param("expense_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	expense, err := h.expenseService.GetUserExpense(ctx, userID, category_id, expenseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ExpenseResponse{
		ID:           expense.ID,
		CategoryID:   expense.CategoryID,
		CategoryName: expense.Category.Name,
		Amount:       expense.Amount,
		Description:  expense.Description,
		Date:         expense.Date,
		CreatedAt:    expense.CreatedAt,
	})

}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	expenses, err := h.expenseService.GetUserExpenses(ctx, category_id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ExpensesListResponse{
		Expenses: expenses,
	})
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	expenseID, err := strconv.Atoi(c.Param("expense_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid expense id",
		})
		return
	}
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}
	if err := h.expenseService.DeleteExpense(ctx, userID, categoryID, expenseID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "expense deleted successfully",
	})

}

func (h *ExpenseHandler) GetAnalytics(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var period dto.ExpensePeriod
	if err := c.BindJSON(&period); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}

	analytics, err := h.expenseService.GetExpenseAnalytics(ctx, userID, categoryID, period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ExpenseAnalytics{
		Period:               analytics.Period,
		TotalAmount:          analytics.TotalAmount,
		ExpensesCount:        analytics.ExpensesCount,
		AveragePerDay:        analytics.AveragePerDay,
		LargestExpense:       analytics.LargestExpense,
		SmallestExpense:      analytics.SmallestExpense,
		AverageExpenseAmount: analytics.AverageExpenseAmount,
	})

}
