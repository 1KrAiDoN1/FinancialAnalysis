package middleware

import (
	"context"
	"errors"
	"finance/internal/dto"
	"finance/internal/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService services.AuthServiceInterface) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Извлекаем токен из заголовка "Bearer TOKEN"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		req := dto.AccessTokenRequest{
			AccessToken: tokenParts[1],
		}

		// Валидация токена через сервис
		userID, err := authService.ValidateToken(ctx, req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Сохраняем userID в контексте для использования в handlers
		c.Set("userID", userID)
		c.Next()
	})
}

func GetUserId(c *gin.Context) (uint, error) {
	userID, ok := c.Get("userID")

	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userID.(uint), nil
}
