package dto

import "time"

// UserInfo - краткая информация о пользователе для ответа
type UserInfo struct {
	// ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type UserID struct {
	UserID int `json:"id"`
}

// Профиль пользователя

// UserProfile - полная информация профиля
type UserProfile struct {
	// ID              uint       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

// ChangePasswordRequest - смена пароля
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}

// UserStats - статистика пользователя
type UserStats struct {
	TotalExpenses   float64           `json:"total_expenses"`
	TotalCategories int               `json:"total_categories"`
	TotalBudgets    int               `json:"total_budgets"`
	MonthlyExpenses float64           `json:"monthly_expenses"`
	WeeklyExpenses  float64           `json:"weekly_expenses"`
	TopCategories   []CategoryExpense `json:"top_categories"`
}
