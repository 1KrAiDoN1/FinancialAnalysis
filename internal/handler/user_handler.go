package handler

import (
	"context"
	"finance/internal/dto"
	"finance/internal/middleware"
	"finance/internal/services"
	"finance/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile godoc
// @Summary Получение профиля пользователя
// @Description Получение информации о текущем пользователе
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserProfile "Профиль пользователя"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /user/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	log := logger.New("user_handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		log.Error("getting user_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	profile, err := h.userService.GetProfile(ctx, userID)
	if err != nil {
		log.Error("getting user profile failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("getting user profile succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.UserProfile{
		Email:     profile.Email,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		CreatedAt: profile.CreatedAt,
	})
}

// DeleteAccount godoc
// @Summary Удаление аккаунта пользователя
// @Description Полное удаление аккаунта пользователя и всех связанных данных
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string "Аккаунт успешно удален"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /user/account [delete]
func (h *UserHandler) DeleteAccount(c *gin.Context) { // нужно придумать, как диактивировать токены, которые были выданы
	log := logger.New("user_handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		log.Error("getting user_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err = h.userService.DeleteAccount(ctx, userID)
	if err != nil {
		log.Error("deleting user account failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("deleting user account succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})

}

// GetStats godoc
// @Summary Получение статистики пользователя
// @Description Получение общей статистики по расходам, категориям и бюджетам
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserStats "Статистика пользователя"
// @Failure 401 {object} dto.ErrorResponse "Требуется авторизация"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /user/stats [get]
func (h *UserHandler) GetStats(c *gin.Context) {
	log := logger.New("user_handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	userID, err := middleware.GetUserId(c)
	if err != nil {
		log.Error("getting user_id failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	stats, err := h.userService.GetUserStats(ctx, userID)
	if err != nil {
		log.Error("getting user stats failed", map[string]interface{}{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("getting user stats succeed", map[string]interface{}{
		"status": http.StatusOK,
	})
	c.JSON(http.StatusOK, dto.UserStats{
		TotalExpenses:   stats.TotalExpenses,
		TotalCategories: stats.TotalCategories,
		TotalBudgets:    stats.TotalBudgets,
		MonthlyExpenses: stats.MonthlyExpenses,
		WeeklyExpenses:  stats.WeeklyExpenses,
	})

}
