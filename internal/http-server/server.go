package server

import (
	//"context"
	"context"
	"finance/internal/config"
	"finance/internal/container"
	"finance/internal/middleware"
	"finance/internal/routes"
	"finance/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	container *container.Container
	router    *gin.Engine
}

func NewServer(container *container.Container) *Server {
	router := gin.Default()

	return &Server{
		container: container,
		router:    router,
	}
}

func (s *Server) Run() error {
	log := logger.New("http-server", true)
	s.setupRoutes()
	serverPort, err := config.LoadConfigServer("./internal/config/config.yaml")
	if err != nil {
		return err
	}

	// Канал для ошибок сервера
	serverErr := make(chan error, 1)
	go func() {
		log.Info("Starting server", map[string]interface{}{
			"port": serverPort.Port,
		})
		if err := s.router.Run(":" + serverPort.Port); err != nil {
			serverErr <- fmt.Errorf("server error: %w", err)
		}
		close(serverErr)
	}()

	// Ожидаем сигналы завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Блокируем до получения сигнала или ошибки сервера
	select {
	case err := <-serverErr:
		return err
	case sig := <-quit:
		log.Info("Shutting down...", map[string]interface{}{
			"Received signal": sig,
		})

		// Получаем доступ к внутреннему http.Server
		srv := &http.Server{
			Addr:    ":" + serverPort.Port,
			Handler: s.router,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Error("server shutdown failed", map[string]interface{}{
				"error": logger.PrettyPrint(err),
			})
		}
		log.Info("Server gracefully stopped", nil)
		return nil
	}

}

func (s *Server) setupRoutes() {
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := s.router.Group("/api/v1")

	// Public routes
	routes.SetupAuthRoutes(api, s.container.Handlers.AuthHandlerInterface)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(s.container.Services.AuthServiceInterface))
	{
		routes.SetupUserRoutes(protected, s.container.Handlers.UserHandlerInterface)
		routes.SetupCategoryRoutes(protected, s.container.Handlers.CategoryHandlerInterface)
		routes.SetupExpenseRoutes(protected, s.container.Handlers.ExpenseHandlerInterface)
		routes.SetupBudgetRoutes(protected, s.container.Handlers.BudgetHandlerInterface)
	}
}
