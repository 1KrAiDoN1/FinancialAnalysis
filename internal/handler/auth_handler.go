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

type AuthHandler struct {
	authService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// SignUp godoc
// @Summary Регистрация нового пользователя
// @Description Создание нового аккаунта пользователя в системе
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "Данные для регистрации"
// @Success 201 {object} dto.UserInfo "Пользователь успешно зарегистрирован"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 409 {object} dto.ErrorResponse "Пользователь с таким email уже существует"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/sign-up [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	log := logger.New("auth-handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var userReg dto.RegisterRequest
	if err := c.BindJSON(&userReg); err != nil {
		log.Error("Invalid Register request", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.SignUp(ctx, userReg)
	if err != nil {
		log.Error("Register failed", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("User registered", map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	})
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// SignIn godoc
// @Summary Вход в систему
// @Description Аутентификация пользователя и получение JWT токенов
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Данные для входа"
// @Success 200 {object} dto.AuthResponse "Успешная аутентификация"
// @Failure 400 {object} dto.ErrorResponse "Неверные данные для входа"
// @Failure 401 {object} dto.ErrorResponse "Неверный email или пароль"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	log := logger.New("auth-handler", true)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var userAuth dto.LoginRequest
	if err := c.BindJSON(&userAuth); err != nil {
		log.Error("Invalid login request", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.SignIn(ctx, userAuth)
	if err != nil {
		log.Error("Login failed", map[string]interface{}{
			"error":  err.Error(),
			"email":  userAuth.Email,
			"status": http.StatusUnauthorized,
		})
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	middleware.SetRefreshTokenCookie(c, token.RefreshToken)

	c.Header("Authorization", "Bearer "+token.AccessToken)

	log.Info("User logged in", map[string]interface{}{
		"user_id": token.User.ID,
		"email":   token.User.Email,
	})
	c.JSON(http.StatusOK, dto.AuthResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User: dto.UserInfo{
			ID:        token.User.ID,
			Email:     token.User.Email,
			FirstName: token.User.FirstName,
			LastName:  token.User.LastName,
		},
	},
	)
}

// Logout godoc
// @Summary Выход из системы
// @Description Деактивация refresh токена и выход из системы
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Успешный выход"
// @Failure 401 {object} dto.ErrorResponse "Токен не найден"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) { //нужно придумать, как диактивировать access-токены, которые были выданы
	log := logger.New("auth-handler", true)
	refresh_token, err := c.Cookie("refresh_token")
	if err != nil {
		log.Error("Logout failed", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(
		"refresh_token",
		refresh_token,
		-1,
		"/",
		"",
		true,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})

}
