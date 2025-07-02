package repositories

import (
	"context"
	"finance/internal/models"
	storage "finance/internal/storages"
	"fmt"
	"time"
)

type ExpenseRepository struct {
	storage storage.ExpenseStorageInterface
}

func NewExpenseRepository(storage storage.ExpenseStorageInterface) *ExpenseRepository { //конструктор
	return &ExpenseRepository{
		storage: storage,
	}
}

func (e *ExpenseRepository) CreateExpense(ctx context.Context, expense models.Expense) (models.Expense, error) {
	query := `INSERT INTO expenses (user_id, category_id, amount, description, date, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, category_id, amount, description, date, created_at`
	result, err := e.storage.CreateExpense(ctx, query, expense)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil
}

func (e *ExpenseRepository) GetExpenseByID(ctx context.Context, userID uint, category_id int, id uint) (models.Expense, error) {
	query := `SELECT e.id, e.user_id, e.category_id, c.name as category_name, e.amount, e.description, e.date, e.created_at FROM expenses e JOIN categories c ON e.category_id = c.id WHERE e.id = $1 AND e.user_id = $2 AND e.category_id = $3`
	result, err := e.storage.GetExpenseByID(ctx, query, id, category_id, userID)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil
}

func (e *ExpenseRepository) GetExpensesByUserID(ctx context.Context, category_id int, userID uint) ([]models.Expense, error) {
	query := `SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
		       e.amount, e.description, e.date, e.created_at
		FROM expenses e
		JOIN categories c ON e.category_id = c.id
		WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
		ORDER BY e.date DESC
	`
	result, err := e.storage.GetExpensesByUserID(ctx, query, category_id, userID)
	if err != nil {
		return []models.Expense{}, err
	}
	return result, nil
}

func (e *ExpenseRepository) GetExpensesByPeriod(ctx context.Context, userID uint, category_id int, period string) ([]models.Expense, error) {
	var query string
	switch period {
	case "weekly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 week'
			ORDER BY e.date DESC
		`
	case "monthly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 month'
			ORDER BY e.date DESC
		`
	case "yearly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 year'
			ORDER BY e.date DESC
		`
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}
	result, err := e.storage.GetExpensesByUserID(ctx, query, category_id, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *ExpenseRepository) DeleteExpense(ctx context.Context, userID uint, category_id int, id uint) error {
	query := `DELETE FROM expenses WHERE id = $1 AND user_id = $2 AND category_id = $3`
	err := e.storage.DeleteExpense(ctx, query, userID, category_id, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExpenseRepository) DeleteExpensesInCategory(ctx context.Context, userID uint, categoryID int) error {
	query := `DELETE FROM expenses WHERE user_id = $1 AND category_id = $2`
	err := e.storage.DeleteExpensesInCategory(ctx, query, userID, categoryID)
	if err != nil {
		return err
	}
	return nil

}

func (e *ExpenseRepository) GetExpensesByCategory(ctx context.Context, userID uint, categoryID int) ([]models.Expense, error) {
	query := `
		SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
		       e.amount, e.description, e.date, e.created_at
		FROM expenses e
		JOIN categories c ON e.category_id = c.id
		WHERE e.user_id = $1 AND e.category_id = $2
		ORDER BY e.date DESC
	`
	result, err := e.storage.GetExpensesByCategory(ctx, query, userID, categoryID)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (e *ExpenseRepository) GetLargestExpenseByPeriod(ctx context.Context, userID uint, category_id int, period string) (models.Expense, error) {
	var query string

	switch period {
	case "weekly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 week'
			ORDER BY e.amount DESC
			LIMIT 1
		`
	case "monthly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 month'
			ORDER BY e.amount DESC
			LIMIT 1
		`
	case "yearly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 year'
			ORDER BY e.amount DESC
			LIMIT 1
		`
	default:
		return models.Expense{}, fmt.Errorf("invalid period: %s", period)
	}
	result, err := e.storage.GetLargestExpenseByPeriod(ctx, query, userID, category_id, period)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil

}

func (e *ExpenseRepository) GetSmallestExpenseByPeriod(ctx context.Context, userID uint, category_id int, period string) (models.Expense, error) {
	var query string

	switch period {
	case "weekly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 week'
			ORDER BY e.amount ASC
			LIMIT 1
		`
	case "monthly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 month'
			ORDER BY e.amount ASC
			LIMIT 1
		`
	case "yearly":
		query = `
			SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
			       e.amount, e.description, e.date, e.created_at
			FROM expenses e
			JOIN categories c ON e.category_id = c.id
			WHERE e.user_id = $1 AND ($2 = 0 OR e.category_id = $2)
			  AND e.date >= NOW() - INTERVAL '1 year'
			ORDER BY e.amount ASC
			LIMIT 1
		`
	default:
		return models.Expense{}, fmt.Errorf("invalid period: %s", period)
	}
	result, err := e.storage.GetSmallestExpenseByPeriod(ctx, query, userID, category_id, period)
	if err != nil {
		return models.Expense{}, err
	}
	return result, nil
}

func (e *ExpenseRepository) GetExpensesByCategoryAndPeriod(ctx context.Context, userID uint, categoryID int, startDate, endDate time.Time) ([]models.Expense, error) {
	query := `
		SELECT e.id, e.user_id, e.category_id, c.name as category_name, 
		       e.amount, e.description, e.date, e.created_at
		FROM expenses e
		JOIN categories c ON e.category_id = c.id
		WHERE e.user_id = $1 AND e.category_id = $2
		  AND e.date >= $3 AND e.date <= $4
		ORDER BY e.date DESC
	`
	result, err := e.storage.GetExpensesByCategoryAndPeriod(ctx, query, userID, categoryID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return result, nil

}
