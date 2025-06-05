package models

import (
	"time"
)

type Expense struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Budget struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	CategoryID uint      `json:"category_id" gorm:"not null"`
	Amount     float64   `json:"amount" gorm:"not null"`
	Period     string    `json:"period"` // monthly, weekly, yearly
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint   `json:"user_id" gorm:"not null;index"`
	Name        string `json:"name" gorm:"not null;size:100"`
	Description string `json:"description" gorm:"size:500"`
	Color       string `json:"color" gorm:"size:7;default:'#3B82F6'"` // Hex color code
	Icon        string `json:"icon" gorm:"size:50;default:'💰'"`       // Emoji or icon name

	// Category settings
	IsActive  bool `json:"is_active" gorm:"default:true"`
	IsDefault bool `json:"is_default" gorm:"default:false"`
	SortOrder int  `json:"sort_order" gorm:"default:0"`

	// Budget settings for this category
	MonthlyBudget *float64 `json:"monthly_budget" gorm:"type:decimal(15,2)"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`

	// Relationships
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Expenses []Expense `json:"expenses,omitempty" gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL"`
	Budgets  []Budget  `json:"budgets,omitempty" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}

type RefreshToken struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id" gorm:"not null;index"`
	Token     string `json:"token" gorm:"uniqueIndex;not null;size:255"`
	IsRevoked bool   `json:"is_revoked" gorm:"default:false"`

	// Token metadata
	DeviceInfo string `json:"device_info" gorm:"size:255"`
	IPAddress  string `json:"ip_address" gorm:"size:45"` // IPv6 support
	UserAgent  string `json:"user_agent" gorm:"size:500"`

	// Token timing
	ExpiresAt  time.Time  `json:"expires_at" gorm:"not null;index"`
	LastUsedAt *time.Time `json:"last_used_at"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type AccessToken struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id" gorm:"not null;index"`
	Token     string `json:"token" gorm:"uniqueIndex;not null;size:255"`
	JTI       string `json:"jti" gorm:"uniqueIndex;not null;size:36"` // JWT ID
	IsRevoked bool   `json:"is_revoked" gorm:"default:false"`

	// Token metadata
	DeviceInfo string `json:"device_info" gorm:"size:255"`
	IPAddress  string `json:"ip_address" gorm:"size:45"`
	UserAgent  string `json:"user_agent" gorm:"size:500"`

	// Token timing
	ExpiresAt time.Time `json:"expires_at" gorm:"not null;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
