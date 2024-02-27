package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Logger interface {
	Debug(...any)
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

// New creates postgres database instance
func New(ctx context.Context, log Logger, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	log.Debug("opened DB")

	err = migrate(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}
	log.Debug("performed DB migration")

	return db, nil
}

func migrate(ctx context.Context, db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.UpContext(ctx, db, "migrations"); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
