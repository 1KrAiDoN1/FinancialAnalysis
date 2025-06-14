package handler

import "finance/internal/services"

type Handlers struct {
	AuthHandlerInterface
	BudgetHandlerInterface
	CategoryHandlerInterface
	ExpenseHandlerInterface
	UserHandlerInterface
}

func NewHandlers(service *services.Services) *Handlers {
	return &Handlers{
		AuthHandlerInterface:     NewAuthHandler(service.AuthServiceInterface),
		BudgetHandlerInterface:   NewBudgetHandler(service.BudgetServiceInterface),
		CategoryHandlerInterface: NewCategoryHandler(service.CategoryServiceInterface),
		ExpenseHandlerInterface:  NewExpenseHandler(service.ExpenseServiceInterface),
		UserHandlerInterface:     NewUserHandler(service.UserServiceInterface),
	}
}
