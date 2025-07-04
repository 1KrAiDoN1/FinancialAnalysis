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

type CategoryHandler struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryHandler(categoryService services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory godoc
// @Summary Создание новой категории
// @Description Создание новой категории расходов для пользователя
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category body dto.CreateCategoryRequest true "Данные для создания категории"
// @Success 200 {object} dto.CategoryResponse "Категория успешно создана"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	log := logger.New("category_handler", true)
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
	var category dto.CreateCategoryRequest
	if err := c.BindJSON(&category); err != nil {
		log.Error("parsing JSON failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newcategory, err := h.categoryService.CreateCategory(ctx, userID, category)
	if err != nil {
		log.Error("creating category failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("creating category succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.CategoryResponse{
		ID:            newcategory.ID,
		Name:          newcategory.Name,
		CreatedAt:     newcategory.CreatedAt,
		ExpensesCount: 0,
		TotalAmount:   0,
	})
}

// GetCategoryByID godoc
// @Summary Получение категории по ID
// @Description Получение информации о конкретной категории пользователя
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category_id path int true "ID категории"
// @Success 200 {object} dto.CategoryResponse "Информация о категории"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID категории"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 404 {object} dto.ErrorResponse "Категория не найдена"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/{category_id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	log := logger.New("category_handler", true)
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
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		log.Error("getting category_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}
	category, err := h.categoryService.GetCategoryByID(ctx, userID, categoryID)
	if err != nil {
		log.Error("getting category failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusNotFound,
		})
		c.JSON(http.StatusNotFound, gin.H{
			"error": "category not found",
		})
		return
	}
	c.JSON(http.StatusOK, dto.CategoryResponse{
		ID:            category.ID,
		Name:          category.Name,
		CreatedAt:     category.CreatedAt,
		ExpensesCount: category.ExpensesCount,
		TotalAmount:   category.TotalAmount,
	})
}

// GetCategories godoc
// @Summary Получение списка категорий
// @Description Получение всех категорий пользователя
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.CategoriesListResponse "Список категорий"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	log := logger.New("category_handler", true)
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
	categories, err := h.categoryService.GetUserCategories(ctx, userID)
	if err != nil {
		log.Error("getting categories failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("getting user categories succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.CategoriesListResponse{
		Categories: categories,
	})
}

// GetMostUsedCategories godoc
// @Summary Получение наиболее используемых категорий
// @Description Получение списка категорий, отсортированных по частоте использования
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.CategoriesListResponse "Список наиболее используемых категорий"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/top [get]
func (h *CategoryHandler) GetMostUsedCategories(c *gin.Context) {
	log := logger.New("category_handler", true)
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
	categories, err := h.categoryService.GetMostUsedCategories(ctx, userID)
	if err != nil {
		log.Error("getting most used categories failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("getting most used categories succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.CategoriesListResponse{
		Categories: categories,
	})

}

// DeleteCategory godoc
// @Summary Удаление категории
// @Description Удаление категории и всех связанных с ней расходов
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category_id path int true "ID категории"
// @Success 200 {object} map[string]string "Категория успешно удалена"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID категории"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/{category_id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	log := logger.New("category_handler", true)
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
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		log.Error("getting category_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}
	err = h.categoryService.DeleteCategory(ctx, userID, categoryID)
	if err != nil {
		log.Error("deleting category failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("deleting category succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted",
	})

}

// GetAnalyticsByCategory godoc
// @Summary Получение аналитики по категории
// @Description Получение детальной аналитики расходов по конкретной категории
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category_id path int true "ID категории"
// @Param period body dto.CategoryPeriod true "Период для анализа"
// @Success 200 {object} dto.CategoryAnalytics "Аналитика по категории"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID категории или период"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /categories/{category_id}/analytics [get]
func (h *CategoryHandler) GetAnalyticsByCategory(c *gin.Context) {
	log := logger.New("category_handler", true)
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
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		log.Error("getting category_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}
	var period dto.CategoryPeriod
	if err := c.BindJSON(&period); err != nil {
		log.Error("parsing JSON failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	category_analytics, err := h.categoryService.GetAnalyticsByCategory(ctx, userID, categoryID, period)
	if err != nil {
		log.Error("getting category analytics failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("getting category analytics succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.CategoryAnalytics{
		CategoryID:           category_analytics.CategoryID,
		CategoryName:         category_analytics.CategoryName,
		Period:               category_analytics.Period,
		TotalAmount:          category_analytics.TotalAmount,
		ExpensesCount:        category_analytics.ExpensesCount,
		AveragePerDay:        category_analytics.AveragePerDay,
		AverageExpenseAmount: category_analytics.AverageExpenseAmount,
		LargestExpense:       category_analytics.LargestExpense,
		SmallestExpense:      category_analytics.SmallestExpense,
	})
}
