package main

import (
	"finance/internal/container"
	server "finance/internal/http-server"
	"log"
)

func main() {

	container, err := container.NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container:", err)
	}
	defer container.Close()

	// Initialize and start server
	srv := server.NewServer(container)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
