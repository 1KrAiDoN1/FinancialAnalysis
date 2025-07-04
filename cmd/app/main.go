package main

import (
	"finance/internal/container"
	server "finance/internal/http-server"
	"finance/pkg/logger"
)

func main() {
	log := logger.New("finance-service", true)

	container, err := container.NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container", map[string]interface{}{
			"error": err.Error(),
			"stack": logger.PrettyPrint(err),
		})
	}
	defer container.Close()

	// Initialize and start server
	srv := server.NewServer(container)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start server", map[string]interface{}{
			"error": err.Error(),
			"stack": logger.PrettyPrint(err),
		})
	}

}
