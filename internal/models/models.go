package models

import (
	"time"
)

type Expense struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	CategoryID   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Amount       float64   `json:"amount"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Budget struct {
	ID          uint      `json:"budget_id"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Amount      float64   `json:"amount"`
	SpentAmount float64   `json:"spent_amount"`
	Period      string    `json:"period"` // monthly, weekly, yearly
	StartDate   time.Time `json:"start_date,omitempty"`
	EndDate     time.Time `json:"end_date,omitempty"`
	//CreatedAt   time.Time `json:"created_at"`
}

type Category struct {
	ID     uint   `json:"category_id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"category_name"`
	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	// Relationships
	User     User      `json:"user,omitempty"`
	Expenses []Expense `json:"expenses,omitempty"`
	Budgets  []Budget  `json:"budgets,omitempty"`
}

type AccessToken struct {
	Token string `json:"access_token"`
	// Token timing
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type RefreshToken struct {
	Token     string    `json:"refresh_token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type UserStats struct {
	TotalExpenses   float64    `json:"total_expenses"`
	TotalCategories int        `json:"total_categories"`
	TotalBudgets    int        `json:"total_budgets"`
	MonthlyExpenses float64    `json:"monthly_expenses"`
	WeeklyExpenses  float64    `json:"weekly_expenses"`
	TopCategories   []Category `json:"categories"`
}
