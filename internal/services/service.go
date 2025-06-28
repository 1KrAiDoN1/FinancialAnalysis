package services

import "finance/internal/repositories"

type Services struct { // создаем структуру, которая будет содержать интерфейсы
	AuthServiceInterface
	ExpenseServiceInterface
	CategoryServiceInterface
	UserServiceInterface
	BudgetServiceInterface
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		AuthServiceInterface:     NewAuthService(repo.AuthRepositoryInterface),
		BudgetServiceInterface:   NewBudgetService(repo.BudgetRepositoryInterface, repo.ExpenseRepositoryInterface),
		ExpenseServiceInterface:  NewExpenseService(repo.ExpenseRepositoryInterface, repo.BudgetRepositoryInterface),
		CategoryServiceInterface: NewCategoryService(repo.CategoryRepositoryInterface),
		UserServiceInterface:     NewUserService(repo.UserRepositoryInterface),
	}

}
