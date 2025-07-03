package dto

import "time"

// CategoryExpense - расходы по категории
type CategoryExpense struct {
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Amount       float64 `json:"amount"`
	Percentage   float64 `json:"percentage"`
}

type CategoryAnalytics struct {
	CategoryID           uint            `json:"category_id"`
	CategoryName         string          `json:"category_name"`
	Period               string          `json:"period"`
	TotalAmount          float64         `json:"total_amount"`
	ExpensesCount        int             `json:"expenses_count"`
	AveragePerDay        float64         `json:"average_per_day"`
	AverageExpenseAmount float64         `json:"average_expense_amount"`
	LargestExpense       ExpenseResponse `json:"largest_expense"`
	SmallestExpense      ExpenseResponse `json:"smallest_expense"`
}

type CategoryPeriod struct {
	Period string `json:"period"`
}

// Запросы для категорий

// CreateCategoryRequest - создание категории
type CreateCategoryRequest struct {
	Name string `json:"category_name" validate:"required,min=1,max=100"`
}

// Ответы для категорий

// CategoryResponse - информация о категории
type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	//Description *string   `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	// Дополнительная информация
	ExpensesCount int     `json:"expenses_count"`
	TotalAmount   float64 `json:"total_amount"`
}

// CategoriesListResponse
type CategoriesListResponse struct {
	Categories []CategoryResponse `json:"categories"`
}
