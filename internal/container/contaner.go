package container

import (
	"context"
	"log"

	//"finance/internal/config"
	"finance/internal/handler"
	"finance/internal/repositories"
	"finance/internal/services"
	storage "finance/internal/storages"
	"finance/internal/storages/database"
)

type Container struct {
	//Config *config.Config
	DB           *database.Storage
	Storages     *storage.Storages
	Repositories *repositories.Repositories
	Services     *services.Services
	Handlers     *handler.Handlers
}

func NewContainer() (*Container, error) {
	ctx := context.Background()
	// Load config
	databaseURL, err := database.NewDatabaseURL()
	if err != nil {
		return nil, err
	}

	// Initialize database
	DB, err := database.NewDatabase(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	dbpool := DB.GetPool()
	storages := storage.NewStorages(dbpool)
	repositories := repositories.NewRepositories(storages)
	services := services.NewServices(repositories)
	handlers := handler.NewHandlers(services)

	return &Container{
		//Config: cfg,
		DB:           DB,
		Storages:     storages,
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		if err := c.DB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}
}
