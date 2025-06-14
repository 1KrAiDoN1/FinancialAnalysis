package main

import (
	"finance/internal/container"
	server "finance/internal/http-server"
	"log"
)

func main() {
	// ctx := context.Background()
	// pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// repository := repository.NewRepository(pool)
	// service := service.NewService(repository)
	// handlers := handler.NewHandler(service)
	// var server server.Server
	// go func() {
	// 	err := server.Run("8080", handlers.InitRoutes())
	// 	if err != nil {
	// 		log.Fatal("Starting Server Error")
	// 	}
	// }()
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit
	// log.Println("Server Stopped")

	// err1 := server.Shutdown(ctx)
	// if err1 != nil {
	// 	log.Fatal("Shutting Down Server Error")
	// }
	// Initialize container with all dependencies
	container, err := container.NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container:", err)
	}

	// Initialize and start server
	srv := server.NewServer(container)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
