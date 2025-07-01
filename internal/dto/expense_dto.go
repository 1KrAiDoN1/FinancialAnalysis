package dto

import "time"

// Запросы для расходов

// CreateExpenseRequest - создание расхода
type CreateExpenseRequest struct {
	CategoryID  uint      `json:"category_id" validate:"required"`
	Amount      float64   `json:"amount" validate:"required,gt=0" example:"25.50"`
	Description string    `json:"description,omitempty" validate:"omitempty,max=500"`
	Date        time.Time `json:"date" validate:"required" example:"2024-01-15T10:30:00Z"`
	Tags        []string  `json:"tags,omitempty" validate:"omitempty,dive,min=1,max=50"`
}

// UpdateExpenseRequest - обновление расхода
type UpdateExpenseRequest struct {
	CategoryID  *uint      `json:"category_id,omitempty" validate:"omitempty"`
	Amount      *float64   `json:"amount,omitempty" validate:"omitempty,gt=0"`
	Description *string    `json:"description,omitempty" validate:"omitempty,max=500"`
	Date        *time.Time `json:"date,omitempty"`
	Tags        *[]string  `json:"tags,omitempty" validate:"omitempty,dive,min=1,max=50"`
}

// Ответы для расходов

// ExpenseResponse - информация о расходе
type ExpenseResponse struct {
	ID           uint             `json:"id"`
	CategoryID   uint             `json:"category_id"`
	CategoryName string           `json:"category_name"`
	Category     CategoryResponse `json:"category,omitempty"`
	Amount       float64          `json:"amount"`
	Description  *string          `json:"description,omitempty"`
	Date         time.Time        `json:"date"`
	CreatedAt    time.Time        `json:"created_at"`
	// UpdatedAt    time.Time        `json:"updated_at"`
}

// ExpensesListResponse - список расходов с пагинацией
type ExpensesListResponse struct {
	Expenses []ExpenseResponse `json:"expenses"`
}

// ExpenseSummary - сводка по расходам
type ExpenseSummary struct {
	TotalAmount   float64 `json:"total_amount"`
	TotalCount    int     `json:"total_count"`
	AverageAmount float64 `json:"average_amount"`
	MinAmount     float64 `json:"min_amount"`
	MaxAmount     float64 `json:"max_amount"`
}

type ExpensePeriod struct {
	Period string `json:"period"`
}

// ExpenseAnalytics - аналитика расходов
type ExpenseAnalytics struct {
	Period               string          `json:"period"`
	TotalAmount          float64         `json:"total_amount"`
	ExpensesCount        int             `json:"expenses_count"`
	AveragePerDay        float64         `json:"average_per_day"`
	LargestExpense       ExpenseResponse `json:"largest_expense"`
	SmallestExpense      ExpenseResponse `json:"smallest_expense"`
	AverageExpenseAmount float64         `json:"average_expense_amount"`
}

// DayExpense - расходы по дням
type DayExpense struct {
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
	Count  int       `json:"count"`
}

// MonthExpense - расходы по месяцам
type MonthExpense struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
}

// ExpenseTrends - тренды расходов
type ExpenseTrends struct {
	GrowthRate     float64 `json:"growth_rate"`      // Процент изменения
	Trend          string  `json:"trend"`            // "increasing", "decreasing", "stable"
	ComparedToPrev float64 `json:"compared_to_prev"` // Сравнение с предыдущим периодом
}
