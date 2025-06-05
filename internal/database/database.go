package database

import (
	"context"
	"finance/internal/config"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, databaseURL string) (*Database, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database connection string: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Проверка соединения
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return &Database{
		pool: pool}, nil

}

func (d *Database) GetPool() *pgxpool.Pool {
	return d.pool
}

func NewDatabaseURL() (string, error) {
	cfg, err := config.SetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to parse config: %w", err)
	}
	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB_username, cfg.DB_password, cfg.DB_host, cfg.DB_port, cfg.DB_name)
	return dbConnStr, nil

}
