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

// List gives listing of user's data
func (r *Repository) List(
	ctx context.Context, user string,
) (entity.DataListing, error) {
	var res entity.DataListing

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT type, name FROM data WHERE login = $1`, user,
	)
	if err != nil {
		return res, fmt.Errorf("failed to select data: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entry entity.DataListEntry
		err = rows.Scan(&entry.Type, &entry.Name)
		if err != nil {
			return res, fmt.Errorf("failed to scan row: %w", err)
		}

		res = append(res, entry)
	}

	err = rows.Err()
	if err != nil {
		return res, fmt.Errorf(
			"failed to finalize data list entries: %w", err)
	}

	return res, nil
}

// Show reveals user's data entry
func (r *Repository) Show(
	ctx context.Context, arg entity.ShowIn,
) (entity.ShowResult, error) {
	var res entity.ShowResult

	err := r.db.QueryRowContext(
		ctx,
		`SELECT type, name, meta, data FROM data WHERE login = $1
		AND type = $2 AND name = $3`, arg.Login, arg.Type, arg.Name,
	).Scan(&res.Type, &res.Name, &res.Meta, &res.Payload)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return res, entity.ErrRecordNotFound
	case err != nil:
		return res, fmt.Errorf("unexpected database error: %w", err)
	}

	return res, nil
}
