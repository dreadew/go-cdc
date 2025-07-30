package main

import (
	"context"
	"log"

	"github.com/dreadew/go-cdc/internal/config"
	"github.com/dreadew/go-cdc/internal/listener"
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

	listener.StartListener(ctx, cfg)
}
