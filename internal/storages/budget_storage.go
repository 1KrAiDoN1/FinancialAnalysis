package storage

import (
	"context"
	"finance/internal/models"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BudgetStorage struct {
	pool *pgxpool.Pool
}

func NewBudgetStorage(pool *pgxpool.Pool) *BudgetStorage {
	return &BudgetStorage{
		pool: pool,
	}
}

func (s *BudgetStorage) CreateBudget(ctx context.Context, query string, budget models.Budget) (models.Budget, error) {
	err := s.pool.QueryRow(ctx, query,
		budget.UserID,
		budget.CategoryID,
		budget.Amount,
		budget.SpentAmount,
		budget.Period,
		budget.StartDate,
		budget.EndDate).Scan(&budget.ID)

	if err != nil {
		return models.Budget{}, fmt.Errorf("failed to create budget: %w", err)
	}

	return budget, nil
}

func (s *BudgetStorage) GetBudgetByID(ctx context.Context, query string, userID uint, category_id int, budget_id int) (models.Budget, error) {
	var budget models.Budget
	row := s.pool.QueryRow(ctx, query, budget_id, userID, category_id)

	err := row.Scan(
		&budget.ID,
		&budget.UserID,
		&budget.CategoryID,
		&budget.Amount,
		&budget.SpentAmount,
		&budget.Period,
		&budget.StartDate,
		&budget.EndDate,
	)

	if err != nil {
		return models.Budget{}, fmt.Errorf("failed to get budget by id: %w", err)
	}

	return budget, nil
}

func (s *BudgetStorage) GetUserBudgets(ctx context.Context, query string, category_id int, userID uint) ([]models.Budget, error) {
	var budgets []models.Budget
	rows, err := s.pool.Query(ctx, query, userID, category_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user budgets: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var budget models.Budget
		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.CategoryID,
			&budget.Amount,
			&budget.SpentAmount,
			&budget.Period,
			&budget.StartDate,
			&budget.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan budget: %w", err)
		}
		budgets = append(budgets, budget)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over budgets: %w", err)
	}

	return budgets, nil
}

func (s *BudgetStorage) DeleteBudget(ctx context.Context, query string, userID uint, category_id int, budget_id int) error {
	result, err := s.pool.Exec(ctx, query, budget_id, userID, category_id)
	if err != nil {
		return fmt.Errorf("failed to delete budget: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("budget not found or access denied")
	}

	return nil
}

func (s *BudgetStorage) DeleteBudgetsInCategory(ctx context.Context, query string, userID uint, categoryID int) error {
	_, err := s.pool.Exec(ctx, query, userID, categoryID)
	if err != nil {
		return fmt.Errorf("failed to delete budgets in category: %w", err)
	}

	return nil
}

func (s *BudgetStorage) UpdateSpentAmount(ctx context.Context, query string, category_id int, budgetID uint, spentAmount float64) error {
	result, err := s.pool.Exec(ctx, query, spentAmount, budgetID, category_id)
	if err != nil {
		return fmt.Errorf("failed to update spent amount: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("budget not found")
	}

	return nil
}

func (s *BudgetStorage) GetActiveBudgetsByCategoryAndDate(ctx context.Context, query string, userID uint, categoryID int, date time.Time) ([]models.Budget, error) {
	var budgets []models.Budget
	rows, err := s.pool.Query(ctx, query, userID, categoryID, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get active budgets by category and date: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var budget models.Budget
		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.CategoryID,
			&budget.Amount,
			&budget.SpentAmount,
			&budget.Period,
			&budget.StartDate,
			&budget.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan budget: %w", err)
		}
		budgets = append(budgets, budget)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over budgets: %w", err)
	}

	return budgets, nil
}
