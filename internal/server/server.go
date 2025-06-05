package server

import (
	"finance/internal/container"
	"finance/internal/middleware"
	"finance/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// type Server struct {
// 	httpServer *http.Server
// }

// func (s *Server) Run(port string, handler http.Handler) error {
// 	s.httpServer = &http.Server{
// 		Addr:    port,
// 		Handler: handler,
// 	}
// 	return s.httpServer.ListenAndServe()
// }

// func (s *Server) Shutdown(ctx context.Context) error {
// 	return s.httpServer.Shutdown(ctx)
// }

type Server struct {
	container *container.Container
	router    *gin.Engine
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
	// viper.AddConfigPath("config")
	// viper.SetConfigName("config")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	return fmt.Errorf("failed to read config: %w", err)
	// }
	// port := viper.GetString("port")
	port := "8081"

	fmt.Printf("Starting server on port %s\n", port)
	return s.router.Run(":" + port)
}

func (s *Server) setupRoutes() {
	api := s.router.Group("/api/v1")

	// Public routes
	routes.SetupAuthRoutes(api, s.container.AuthHandler)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(s.container.AuthService))
	{
		routes.SetupUserRoutes(protected, s.container.UserHandler)
		routes.SetupCategoryRoutes(protected, s.container.CategoryHandler)
		routes.SetupExpenseRoutes(protected, s.container.ExpenseHandler)
		routes.SetupBudgetRoutes(protected, s.container.BudgetHandler)
	}
}
