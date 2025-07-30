package db

import (
	"context"
	"fmt"

	"github.com/dreadew/go-cdc/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := FormatDSN(cfg)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func FormatDSN(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName)
}
