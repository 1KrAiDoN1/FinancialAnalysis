package storage

import (
	"context"
	"finance/internal/models"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseStorage struct {
	pool *pgxpool.Pool
}

func NewExpenseStorage(pool *pgxpool.Pool) *ExpenseStorage {
	return &ExpenseStorage{
		pool: pool,
	}
}

func (s *ExpenseStorage) CreateExpense(ctx context.Context, query string, expense models.Expense) (models.Expense, error) {
	var new_expense models.Expense
	err := s.pool.QueryRow(ctx, query, expense.UserID, expense.CategoryID, expense.Amount, expense.Description, expense.Date, expense.CreatedAt).Scan(&new_expense.ID, &new_expense.CategoryID, &new_expense.Amount, &new_expense.Description, &new_expense.Date, &new_expense.CreatedAt)
	if err != nil {
		return models.Expense{}, err
	}
	return expense, nil
}

func (s *ExpenseStorage) GetExpenseByID(ctx context.Context, query string, userID uint, categoryID int, id uint) (models.Expense, error) {
	var expense models.Expense
	row := s.pool.QueryRow(ctx, query, id, userID, categoryID)

	err := row.Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.CategoryName,
		&expense.Amount,
		&expense.Description,
		&expense.Date,
		&expense.CreatedAt,
	)

	if err != nil {
		return models.Expense{}, fmt.Errorf("failed to get expense by id: %w", err)
	}

	return expense, nil
}

func (s *ExpenseStorage) GetExpensesByUserID(ctx context.Context, query string, categoryID int, userID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	rows, err := s.pool.Query(ctx, query, userID, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses by user id: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.CategoryName,
			&expense.Amount,
			&expense.Description,
			&expense.Date,
			&expense.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %w", err)
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over expenses: %w", err)
	}

	return expenses, nil
}

func (s *ExpenseStorage) GetExpensesByPeriod(ctx context.Context, query string, userID uint, categoryID int, period string) ([]models.Expense, error) {
	var expenses []models.Expense
	rows, err := s.pool.Query(ctx, query, userID, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses by period: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.CategoryName,
			&expense.Amount,
			&expense.Description,
			&expense.Date,
			&expense.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %w", err)
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over expenses: %w", err)
	}

	return expenses, nil
}

func (s *ExpenseStorage) DeleteExpense(ctx context.Context, query string, userID uint, categoryID int, id uint) error {
	result, err := s.pool.Exec(ctx, query, id, userID, categoryID)
	if err != nil {
		return fmt.Errorf("failed to delete expense: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("expense not found or access denied")
	}

	return nil
}

func (s *ExpenseStorage) DeleteExpensesInCategory(ctx context.Context, query string, userID uint, categoryID int) error {
	_, err := s.pool.Exec(ctx, query, userID, categoryID)
	if err != nil {
		return fmt.Errorf("failed to delete expenses in category: %w", err)
	}

	return nil
}

func (s *ExpenseStorage) GetExpensesByCategory(ctx context.Context, query string, userID uint, categoryID int) ([]models.Expense, error) {
	var expenses []models.Expense
	rows, err := s.pool.Query(ctx, query, userID, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses by category: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.CategoryName,
			&expense.Amount,
			&expense.Description,
			&expense.Date,
			&expense.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %w", err)
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over expenses: %w", err)
	}

	return expenses, nil
}

func (s *ExpenseStorage) GetLargestExpenseByPeriod(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error) {
	var expense models.Expense
	row := s.pool.QueryRow(ctx, query, userID, categoryID)

	err := row.Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.CategoryName,
		&expense.Amount,
		&expense.Description,
		&expense.Date,
		&expense.CreatedAt,
	)

	if err != nil {
		return models.Expense{}, fmt.Errorf("failed to get largest expense: %w", err)
	}

	return expense, nil
}

func (s *ExpenseStorage) GetSmallestExpenseByPeriod(ctx context.Context, query string, userID uint, categoryID int, period string) (models.Expense, error) {
	var expense models.Expense
	row := s.pool.QueryRow(ctx, query, userID, categoryID)

	err := row.Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.CategoryName,
		&expense.Amount,
		&expense.Description,
		&expense.Date,
		&expense.CreatedAt,
	)

	if err != nil {
		return models.Expense{}, fmt.Errorf("failed to get smallest expense: %w", err)
	}

	return expense, nil
}

func (s *ExpenseStorage) GetExpensesByCategoryAndPeriod(ctx context.Context, query string, userID uint, categoryID int, startDate, endDate time.Time) ([]models.Expense, error) {
	var expenses []models.Expense
	rows, err := s.pool.Query(ctx, query, userID, categoryID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenses by category and period: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.CategoryName,
			&expense.Amount,
			&expense.Description,
			&expense.Date,
			&expense.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %w", err)
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over expenses: %w", err)
	}

	return expenses, nil
}
