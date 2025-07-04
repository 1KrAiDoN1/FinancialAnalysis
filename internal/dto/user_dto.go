package dto

import "time"

// UserInfo - краткая информация о пользователе для ответа
type UserInfo struct {
	ID        uint   `json:"id" example:"1"`
	Email     string `json:"email" example:"user@example.com"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
}
type UserID struct {
	UserID int `json:"id"`
}

// Профиль пользователя

// UserProfile - полная информация профиля
type UserProfile struct {
	Email     string    `json:"email" example:"user@example.com"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-15T10:30:00Z"`
}

// UserStats структура статистики пользователя
type UserStats struct {
	TotalExpenses   float64 `json:"total_expenses" example:"1250.50"`
	TotalCategories int     `json:"total_categories" example:"5"`
	TotalBudgets    int     `json:"total_budgets" example:"3"`
	MonthlyExpenses float64 `json:"monthly_expenses" example:"450.75"`
	WeeklyExpenses  float64 `json:"weekly_expenses" example:"125.25"`
}

// ChangePasswordRequest - смена пароля
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}
