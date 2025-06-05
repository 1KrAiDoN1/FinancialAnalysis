package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type CategoryRepository struct {
	pool *pgxpool.Pool
}

func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepository { //конструктор
	return &CategoryRepository{
		pool: pool,
	}
}
