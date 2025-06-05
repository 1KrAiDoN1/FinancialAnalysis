package handler

import (
	"finance/internal/services"

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

func (h *UserHandler) GetProfile(c *gin.Context) {
	// ...
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	// ...
}

func (h *UserHandler) GetStats(c *gin.Context) {
	// ...
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
}
