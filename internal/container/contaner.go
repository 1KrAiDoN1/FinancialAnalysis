package container

import (
	"context"
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

	// // Repositories
	// AuthRepo     repositories.AuthRepositoryInterface
	// UserRepo     repositories.UserRepositoryInterface
	// CategoryRepo repositories.CategoryRepositoryInterface
	// ExpenseRepo  repositories.ExpenseRepositoryInterface
	// BudgetRepo   repositories.BudgetRepositoryInterface

	// // Services
	// AuthService     services.AuthServiceInterface
	// UserService     services.UserServiceInterface
	// CategoryService services.CategoryServiceInterface
	// ExpenseService  services.ExpenseServiceInterface
	// BudgetService   services.BudgetServiceInterface

	// // Handlers
	// AuthHandler     handler.AuthHandlerInterface
	// UserHandler     handler.UserHandlerInterface
	// CategoryHandler handler.CategoryHandlerInterface
	// ExpenseHandler  handler.ExpenseHandlerInterface
	// BudgetHandler   handler.BudgetHandlerInterface
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

	// // Initialize repositories
	// authRepo := repositories.NewAuthRepository(dbpool)
	// userRepo := repositories.NewUserRepository(dbpool)
	// categoryRepo := repositories.NewCategoryRepository(dbpool)
	// expenseRepo := repositories.NewExpenseRepository(dbpool)
	// budgetRepo := repositories.NewBudgetRepository(dbpool)

	// // Initialize services
	// authService := services.NewAuthService(authRepo)
	// userService := services.NewUserService(userRepo)
	// categoryService := services.NewCategoryService(categoryRepo)
	// expenseService := services.NewExpenseService(expenseRepo, categoryRepo)
	// budgetService := services.NewBudgetService(budgetRepo, expenseRepo)

	// // Initialize handlers
	// authHandler := handler.NewAuthHandler(authService)
	// userHandler := handler.NewUserHandler(userService)
	// categoryHandler := handler.NewCategoryHandler(categoryService)
	// expenseHandler := handler.NewExpenseHandler(expenseService)
	// budgetHandler := handler.NewBudgetHandler(budgetService)

	return &Container{
		//Config: cfg,
		DB:           DB,
		Storages:     storages,
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,

		// AuthRepo:     authRepo,
		// UserRepo:     userRepo,
		// CategoryRepo: categoryRepo,
		// ExpenseRepo:  expenseRepo,
		// BudgetRepo:   budgetRepo,

		// AuthService:     authService,
		// UserService:     userService,
		// CategoryService: categoryService,
		// ExpenseService:  expenseService,
		// BudgetService:   budgetService,

		// AuthHandler:     authHandler,
		// UserHandler:     userHandler,
		// CategoryHandler: categoryHandler,
		// ExpenseHandler:  expenseHandler,
		// BudgetHandler:   budgetHandler,
	}, nil
}
