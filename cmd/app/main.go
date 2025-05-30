package main

import (
	"context"
	"finance/internal/handler"
	"finance/internal/repository"
	"finance/internal/server"
	"finance/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	pool := &pgxpool.Pool{}
	repository := repository.NewRepository(pool)
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)
	var server server.Server
	go func() {
		err := server.Run("8080", handlers.InitRoutes())
		if err != nil {
			log.Fatal("Starting Server Error")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Server Stopped")

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Shutting Down Server Error")
	}
}
