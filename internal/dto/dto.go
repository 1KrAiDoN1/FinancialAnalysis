package dto

import "time"

// Общие структуры для всех DTO

// PaginationRequest - параметры пагинации
type PaginationRequest struct {
	Page     int `json:"page" validate:"min=1" default:"1"`
	PageSize int `json:"page_size" validate:"min=1,max=100" default:"10"`
}

// PaginationResponse - ответ с пагинацией
type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// ErrorResponse - стандартный формат ошибки
type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

// SuccessResponse - стандартный успешный ответ
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Запросы для аутентификации

// RegisterRequest - данные для регистрации
type RegisterRequest struct {
	Email              string    `json:"email" validate:"required,email" example:"user@example.com"`
	Password           string    `json:"password" validate:"required,min=8,max=100" example:"password123"`
	ConfirmPassword    string    `json:"confirm_password" validate:"required,eqfield=Password"`
	FirstName          string    `json:"first_name" validate:"required,min=2,max=50" example:"John"`
	LastName           string    `json:"last_name" validate:"required,min=2,max=50" example:"Doe"`
	TimeOfRegistration time.Time `json:"time_of_registration"`
}

// LoginRequest - данные для входа
type LoginRequest struct {
	UserID   uint   `json:"id"`
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// RefreshTokenRequest - запрос обновления токена
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type AccessTokenRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Ответы аутентификации

// AuthResponse - ответ после успешной аутентификации
type AuthResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         UserInfo `json:"user"`
}

// UserInfo - краткая информация о пользователе для ответа
type UserInfo struct {
	// ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Профиль пользователя

// UserProfile - полная информация профиля
type UserProfile struct {
	// ID              uint       `json:"id"`
	Email           string     `json:"email"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Avatar          *string    `json:"avatar,omitempty"`
	Currency        string     `json:"currency" example:"USD"`
	Timezone        string     `json:"timezone" example:"UTC"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty"`
	IsEmailVerified bool       `json:"is_email_verified"`
}

// UpdateProfileRequest - обновление профиля
type UpdateProfileRequest struct {
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Avatar    *string `json:"avatar,omitempty" validate:"omitempty,url"`
	Currency  *string `json:"currency,omitempty" validate:"omitempty,len=3"`
	Timezone  *string `json:"timezone,omitempty"`
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
	RecentExpenses  []ExpenseResponse `json:"recent_expenses"`
	BudgetAlerts    []BudgetAlert     `json:"budget_alerts"`
}

// CategoryExpense - расходы по категории
type CategoryExpense struct {
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Amount       float64 `json:"amount"`
	Percentage   float64 `json:"percentage"`
}

// Запросы для категорий

// CreateCategoryRequest - создание категории
type CreateCategoryRequest struct {
	Name string `json:"category_name" validate:"required,min=1,max=100"`
}

// UpdateCategoryRequest - обновление категории
type UpdateCategoryRequest struct {
	Name        *string `json:"category_name,omitempty" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
	Color       *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
	Icon        *string `json:"icon,omitempty" validate:"omitempty,max=50"`
}

// Ответы для категорий

// CategoryResponse - информация о категории
type CategoryResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Color       string    `json:"color"`
	Icon        *string   `json:"icon,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Дополнительная информация
	ExpensesCount int     `json:"expenses_count,omitempty"`
	TotalAmount   float64 `json:"total_amount,omitempty"`
}

// CategoriesListResponse
type CategoriesListResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

// Запросы для расходов

// CreateExpenseRequest - создание расхода
type CreateExpenseRequest struct {
	CategoryID  uint      `json:"category_id" validate:"required"`
	Amount      float64   `json:"amount" validate:"required,gt=0" example:"25.50"`
	Description *string   `json:"description,omitempty" validate:"omitempty,max=500"`
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

// ExpenseFilter - фильтры для поиска расходов
type ExpenseFilter struct {
	PaginationRequest
	CategoryIDs []uint     `json:"category_ids,omitempty" form:"category_ids"`
	MinAmount   *float64   `json:"min_amount,omitempty" form:"min_amount" validate:"omitempty,gte=0"`
	MaxAmount   *float64   `json:"max_amount,omitempty" form:"max_amount" validate:"omitempty,gte=0"`
	DateFrom    *time.Time `json:"date_from,omitempty" form:"date_from"`
	DateTo      *time.Time `json:"date_to,omitempty" form:"date_to"`
	Tags        []string   `json:"tags,omitempty" form:"tags"`
	Search      *string    `json:"search,omitempty" form:"search" validate:"omitempty,min=1,max=100"`
	SortBy      string     `json:"sort_by" form:"sort_by" validate:"omitempty,oneof=date amount category" default:"date"`
	SortOrder   string     `json:"sort_order" form:"sort_order" validate:"omitempty,oneof=asc desc" default:"desc"`
}

// Ответы для расходов

// ExpenseResponse - информация о расходе
type ExpenseResponse struct {
	ID           uint             `json:"id"`
	CategoryID   uint             `json:"category_id"`
	CategoryName string           `json:"category_name"`
	Category     CategoryResponse `json:"category"`
	Amount       float64          `json:"amount"`
	Description  *string          `json:"description,omitempty"`
	Date         time.Time        `json:"date"`
	Tags         []string         `json:"tags,omitempty"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
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
	Period        string            `json:"period"`
	TotalAmount   float64           `json:"total_amount"`
	ExpensesCount int               `json:"expenses_count"`
	AveragePerDay float64           `json:"average_per_day"`
	ByCategory    []CategoryExpense `json:"by_category"`
	ByDay         []DayExpense      `json:"by_day"`
	ByMonth       []MonthExpense    `json:"by_month,omitempty"`
	TopExpenses   []ExpenseResponse `json:"top_expenses"`
	Trends        ExpenseTrends     `json:"trends"`
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

// Запросы для бюджетов

// CreateBudgetRequest - создание бюджета
type CreateBudgetRequest struct {
	CategoryID uint       `json:"category_id" validate:"required"`
	Amount     float64    `json:"amount" validate:"required,gt=0" example:"500.00"`
	Period     string     `json:"period" validate:"required,oneof=weekly monthly yearly" example:"monthly"`
	StartDate  *time.Time `json:"start_date,omitempty"`
	EndDate    *time.Time `json:"end_date,omitempty"`
	IsActive   bool       `json:"is_active" default:"true"`
}

// UpdateBudgetRequest - обновление бюджета
type UpdateBudgetRequest struct {
	Amount    *float64   `json:"amount,omitempty" validate:"omitempty,gt=0"`
	Period    *string    `json:"period,omitempty" validate:"omitempty,oneof=weekly monthly yearly"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	IsActive  *bool      `json:"is_active,omitempty"`
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
