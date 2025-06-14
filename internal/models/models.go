package models

import (
	"time"
)

type Expense struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Budget struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	CategoryID uint      `json:"category_id"`
	Amount     float64   `json:"amount"`
	Period     string    `json:"period"` // monthly, weekly, yearly
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Category struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	// Timestamps
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	// Relationships
	User     User      `json:"user,omitempty"`
	Expenses []Expense `json:"expenses,omitempty"`
	Budgets  []Budget  `json:"budgets,omitempty"`
}

type AccessToken struct {
	Token string `json:"token"`
	// Token timing
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStats struct {
	TotalExpenses   float64    `json:"total_expenses"`
	TotalCategories int        `json:"total_categories"`
	TotalBudgets    int        `json:"total_budgets"`
	MonthlyExpenses float64    `json:"monthly_expenses"`
	WeeklyExpenses  float64    `json:"weekly_expenses"`
	TopCategories   []Category `json:"categories"`
	RecentExpenses  []Expense  `json:"expenses"`
	BudgetAlerts    []Budget   `json:"budget"`
}
