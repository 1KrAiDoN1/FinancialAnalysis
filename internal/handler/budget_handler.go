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
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var budget dto.CreateBudgetRequest
	if err := c.BindJSON(&budget); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newbudget, err := b.budgetService.CreateBudget(ctx, userID, budget)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.BudgetResponse{
		ID:              newbudget.ID,
		CategoryID:      newbudget.CategoryID,
		Amount:          newbudget.Amount,
		CreatedAt:       newbudget.CreatedAt,
		SpentAmount:     newbudget.SpentAmount,
		RemainingAmount: newbudget.Amount - newbudget.SpentAmount,
		Period:          newbudget.Period,
		StartDate:       newbudget.StartDate,
		EndDate:         newbudget.EndDate,
	})
}

func (b *BudgetHandler) GetBudgets(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	budgets, err := b.budgetService.GetUserBudgets(ctx, userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.BudgetsListResponse{
		Budgets: budgets,
	})
}

func (b *BudgetHandler) UpdateBudget(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	budgetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid budget id",
		})
		return
	}
	var upbudget dto.UpdateBudgetRequest
	if err := c.BindJSON(&upbudget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if err := b.budgetService.UpdateBudget(ctx, userID, budgetID, upbudget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "budget updated successfully",
	})

}

func (b *BudgetHandler) DeleteBudget(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	budgetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid budget id",
		})
		return
	}
	if err := b.budgetService.DeleteBudget(ctx, userID, budgetID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "budget deleted successfully",
	})

}
