package dto

// ErrorResponse структура для ошибок
type ErrorResponse struct {
	Error   string            `json:"error" example:"Validation failed"`
	Message string            `json:"message,omitempty" example:"Email is required"`
	Details map[string]string `json:"details,omitempty"`
}
