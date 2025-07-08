package main

import (
	_ "finance/docs"
	"finance/internal/container"
	server "finance/internal/http-server"
	"finance/pkg/logger"
)

// @title Finance API
// @version 1.0
// @description Система управления личными финансами
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	log := logger.New("finance-service", true)

	container, err := container.NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container", map[string]string{
			"error": err.Error(),
			"stack": logger.PrettyPrint(err),
		})
	}
	defer container.Close()

	// Initialize and start server
	srv := server.NewServer(container)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start server", map[string]string{
			"error": err.Error(),
			"stack": logger.PrettyPrint(err),
		})
	}

}
