package models

import "time"

type User struct {
	ID                 string    `json:"id"`
	Username           string    `json:"username" binding:"required"`
	Email              string    `json:"email" binding:"required"`
	Password           string    `json:"password" binding:"required"`
	TimeOfRegistration time.Time `json:"time_of_registration"`
}
