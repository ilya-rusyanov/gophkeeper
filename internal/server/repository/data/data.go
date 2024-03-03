package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
)

// Repository is repository for user's data
type Repository struct {
	db *sql.DB
}

// New constructs new repository
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Store stores user's data
func (r *Repository) Store(
	ctx context.Context, in *entity.StoreIn,
) error {
	var pgErr *pgconn.PgError

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO data (login, type, name, meta, data)
VALUES($1, $2, $3, $4, $5)`,
		in.Login,
		in.Type,
		in.Name,
		in.Meta,
		in.Payload,
	)
	switch {
	case errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation:
		return entity.ErrRecordAlreadyExists
	case err != nil:
		return fmt.Errorf("unexpected database error: %w", err)
	}

	return nil
}
