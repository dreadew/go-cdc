package main

import (
	"context"
	"log"

	"github.com/dreadew/go-cdc/internal/client/db"
	"github.com/dreadew/go-cdc/internal/config"
	"github.com/dreadew/go-cdc/internal/migration"
	"github.com/dreadew/go-cdc/internal/writer"
)

const (
	configPath = "config/config.yaml"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("error while loading configuration: %v", err)
	}

	if err := migration.RungMigrations(cfg); err != nil {
		log.Fatalf("error while running migrations: %v", err)
	}

	pool, err := db.NewPgPool(ctx, cfg)
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer pool.Close()

	writer.StartWriter(ctx, pool)
}
