package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository { //конструктор
	return &UserRepository{
		pool: pool,
	}
}
