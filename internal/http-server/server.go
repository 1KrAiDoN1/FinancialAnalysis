package server

import (
	"context"
	"finance/internal/container"
	"finance/internal/middleware"
	"finance/internal/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	container  *container.Container
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(container *container.Container) *Server {
	router := gin.Default()

	// Global middleware
	// router.Use(middleware.CORS())
	// router.Use(middleware.Logger())
	// router.Use(middleware.ErrorHandler())

	return &Server{
		container: container,
		router:    router,
	}
}

func (s *Server) Run() error {
	s.setupRoutes()
	port := "8081"
	// viper.AddConfigPath("config")
	// viper.SetConfigName("config")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	return fmt.Errorf("failed to read config: %w", err)
	// }
	// port := viper.GetString("port")
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}
	// Канал для ошибок сервера
	serverErr := make(chan error, 1)
	go func() {
		fmt.Printf("Starting server on port %s\n", port)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
		fmt.Printf("Received signal: %s. Shutting down...\n", sig)
	}

	// Настройка graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("forced shutdown: %w", err)
	}

	fmt.Println("Server gracefully stopped")
	return nil
}

func (s *Server) setupRoutes() {
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
