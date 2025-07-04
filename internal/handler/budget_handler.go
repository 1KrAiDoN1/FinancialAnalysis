package handler

import (
	"context"
	"finance/internal/dto"
	"finance/internal/middleware"
	"finance/internal/services"
	"finance/pkg/logger"
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

// CreateBudget godoc
// @Summary Создание нового бюджета
// @Description Создание нового бюджета для указанной категории
// @Tags Budgets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category_id path int true "ID категории"
// @Param budget body dto.CreateBudgetRequest true "Данные для создания бюджета"
// @Success 200 {object} dto.BudgetResponse "Бюджет успешно создан"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/{category_id}/budgets [post]
func (b *BudgetHandler) CreateBudget(c *gin.Context) {
	log := logger.New("budget_handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		log.Error("getting user_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		log.Error("getting category_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	var budget dto.CreateBudgetRequest
	if err := c.BindJSON(&budget); err != nil {
		log.Error("parsing JSON failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newbudget, err := b.budgetService.CreateBudget(ctx, userID, category_id, budget)
	if err != nil {
		log.Error("creating budget failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("creating budget succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.BudgetResponse{
		ID:              newbudget.ID,
		CategoryID:      uint(category_id),
		Amount:          newbudget.Amount,
		CreatedAt:       newbudget.CreatedAt,
		SpentAmount:     newbudget.SpentAmount,
		RemainingAmount: newbudget.Amount - newbudget.SpentAmount,
		Period:          newbudget.Period,
		StartDate:       newbudget.StartDate,
		EndDate:         newbudget.EndDate,
	})
}

// GetBudgets godoc
// @Summary Получение списка бюджетов
// @Description Получение всех бюджетов пользователя для указанной категории
// @Tags Budgets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category_id path int true "ID категории"
// @Success 200 {object} dto.BudgetsListResponse "Список бюджетов"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID категории"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/{category_id}/budgets [get]
func (b *BudgetHandler) GetBudgets(c *gin.Context) {
	log := logger.New("budget_handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		log.Error("getting user_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		log.Error("getting category_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	budgets, err := b.budgetService.GetUserBudgets(ctx, userID, category_id)
	if err != nil {
		log.Error("getting budgets failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("getting budgets succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.BudgetsListResponse{
		Budgets: budgets,
	})
}

// DeleteBudget godoc
// @Summary Удаление бюджета
// @Description Удаление конкретного бюджета пользователя
// @Tags Budgets
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category_id path int true "ID категории"
// @Param budget_id path int true "ID бюджета"
// @Success 200 {object} map[string]string "Бюджет успешно удален"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID категории или бюджета"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/{category_id}/budgets/{budget_id} [delete]
func (b *BudgetHandler) DeleteBudget(c *gin.Context) {
	log := logger.New("budget_handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		log.Error("getting user_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		log.Error("getting category_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	budgetID, err := strconv.Atoi(c.Param("budget_id"))
	if err != nil {
		log.Error("getting budget_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid budget id",
		})
		return
	}
	if err := b.budgetService.DeleteBudget(ctx, userID, category_id, budgetID); err != nil {
		log.Error("deleting budget failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("deleting budget succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "budget deleted successfully",
	})

}
