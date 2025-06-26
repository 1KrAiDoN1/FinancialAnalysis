package dto

import "time"

// Запросы для бюджетов

// CreateBudgetRequest - создание бюджета
type CreateBudgetRequest struct {
	UserID     uint       `json:"user_id"`
	CategoryID uint       `json:"category_id" validate:"required"`
	Amount     float64    `json:"amount" validate:"required,gt=0" example:"500.00"`
	Period     string     `json:"period" validate:"required,oneof=weekly monthly yearly" example:"monthly"`
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	IsActive   bool       `json:"is_active" default:"true"`
}

// UpdateBudgetRequest - обновление бюджета
type UpdateBudgetRequest struct {
	CategoryID uint       `json:"category_id" validate:"required"`
	Amount     *float64   `json:"amount,omitempty" validate:"omitempty,gt=0"`
	Period     *string    `json:"period,omitempty" validate:"omitempty,oneof=weekly monthly yearly"`
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	IsActive   *bool      `json:"is_active,omitempty"`
}

// Ответы для бюджетов

// BudgetResponse - информация о бюджете
type BudgetResponse struct {
	ID         uint             `json:"id"`
	CategoryID uint             `json:"category_id"`
	Category   CategoryResponse `json:"category"`
	Amount     float64          `json:"amount"`
	Period     string           `json:"period"`
	StartDate  *time.Time       `json:"start_date,omitempty"`
	EndDate    *time.Time       `json:"end_date,omitempty"`
	IsActive   bool             `json:"is_active"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`

	// Статистика
	SpentAmount     float64 `json:"spent_amount"`
	RemainingAmount float64 `json:"remaining_amount"`
	SpentPercentage float64 `json:"spent_percentage"`
	DaysRemaining   int     `json:"days_remaining,omitempty"`
}

// BudgetsListResponse - список бюджетов
type BudgetsListResponse struct {
	Budgets []BudgetResponse `json:"budgets"`
}

// BudgetStatus - статус бюджета
type BudgetStatus struct {
	BudgetID        uint    `json:"budget_id"`
	CategoryName    string  `json:"category_name"`
	BudgetAmount    float64 `json:"budget_amount"`
	SpentAmount     float64 `json:"spent_amount"`
	RemainingAmount float64 `json:"remaining_amount"`
	SpentPercentage float64 `json:"spent_percentage"`
	Status          string  `json:"status"` // "ok", "warning", "exceeded"
	DaysRemaining   int     `json:"days_remaining"`
}

// BudgetAlert - уведомление о бюджете
type BudgetAlert struct {
	BudgetID     uint      `json:"budget_id"`
	CategoryName string    `json:"category_name"`
	AlertType    string    `json:"alert_type"` // "warning", "exceeded", "near_end"
	Message      string    `json:"message"`
	Percentage   float64   `json:"percentage"`
	CreatedAt    time.Time `json:"created_at"`
}
