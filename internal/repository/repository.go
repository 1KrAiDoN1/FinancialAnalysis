package repository

import (
	"finance/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(pool),
	}
}
