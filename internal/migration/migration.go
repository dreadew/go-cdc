package migration

import (
	"github.com/dreadew/go-cdc/internal/client/db"
	"github.com/dreadew/go-cdc/internal/config"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RungMigrations(cfg *config.Config) error {
	dsn := db.FormatDSN(cfg)
	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
