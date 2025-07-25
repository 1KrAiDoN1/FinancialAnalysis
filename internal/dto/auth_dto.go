package dto

import "time"

// Запросы для аутентификации

// RegisterRequest - данные для регистрации
type RegisterRequest struct {
	Email              string    `json:"email" validate:"required,email" example:"user@example.com"`
	Password           string    `json:"password" validate:"required,min=8,max=100" example:"password123"`
	ConfirmPassword    string    `json:"confirm_password" validate:"required,eqfield=Password"`
	FirstName          string    `json:"first_name" validate:"required,min=2,max=50" example:"John"`
	LastName           string    `json:"last_name" validate:"required,min=2,max=50" example:"Doe"`
	TimeOfRegistration time.Time `json:"time_of_registration" example:"2024-01-15T10:30:00Z"`
}

// LoginRequest - данные для входа
type LoginRequest struct {
	// UserID   uint   `json:"id"`
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// RefreshTokenRequest - запрос обновления токена
type RefreshTokenRequest struct {
	RefreshToken string    `json:"refresh_token" validate:"required"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type AccessTokenRequest struct {
	AccessToken string    `json:"access_token" validate:"required"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Ответы аутентификации

// AuthResponse - ответ после успешной аутентификации
type AuthResponse struct {
	AccessToken  string   `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string   `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User         UserInfo `json:"user"`
}
