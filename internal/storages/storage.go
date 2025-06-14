package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storages struct {
	AuthStorageInterface
	BudgetStorageInterface
	CategoryStorageInterface
	ExpenseStorageInterface
	UserStorageInterface
}

func NewStorages(pool *pgxpool.Pool) *Storages {
	return &Storages{
		AuthStorageInterface:     NewAuthStorage(pool),
		BudgetStorageInterface:   NewBudgetStorage(pool),
		CategoryStorageInterface: NewCategoryStorage(pool),
		ExpenseStorageInterface:  NewExpenseStorage(pool),
		UserStorageInterface:     NewUserStorage(pool),
	}
}
