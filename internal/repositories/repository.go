package repositories

import storage "finance/internal/storages"

type Repositories struct {
	AuthRepositoryInterface
	BudgetRepositoryInterface
	CategoryRepositoryInterface
	ExpenseRepositoryInterface
	UserRepositoryInterface
}

func NewRepositories(storage *storage.Storages) *Repositories {
	return &Repositories{
		AuthRepositoryInterface:     NewAuthRepository(storage.AuthStorageInterface),
		BudgetRepositoryInterface:   NewBudgetRepository(storage.BudgetStorageInterface),
		CategoryRepositoryInterface: NewCategoryRepository(storage.CategoryStorageInterface),
		ExpenseRepositoryInterface:  NewExpenseRepository(storage.ExpenseStorageInterface),
		UserRepositoryInterface:     NewUserRepository(storage.UserStorageInterface),
	}
}
