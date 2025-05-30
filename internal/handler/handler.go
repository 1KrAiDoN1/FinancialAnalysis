package handler

import (
	"finance/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")

	}
	// api := router.Group("/api")
	// {

	// }
	return router
}
