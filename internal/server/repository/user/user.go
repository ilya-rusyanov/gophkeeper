package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
)

// Repo is repository for manipulating with users and their credentials
type Repo struct {
	db *sql.DB
}

// New constructs user repository
func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// Store stores user credentials
func (r *Repo) Store(
	ctx context.Context, creds entity.UserCredentials,
) error {
	var pgErr *pgconn.PgError

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users (login, password)
VALUES($1, $2)`, creds.Login, creds.Password)
	switch {
	case errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation:
		return entity.ErrUserAlreadyExists
	case err != nil:
		return fmt.Errorf("unexpected database error: %w", err)
	}

	return nil
}

// GetByUserName returns password by user's name
func (r *Repo) GetByUsername(
	ctx context.Context, login string,
) (string, error) {
	var res string

	row := r.db.QueryRowContext(
		ctx,
		`SELECT password FROM users WHERE login = $1`,
		login,
	)
	err := row.Scan(&res)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return res, entity.ErrNoSuchUser
	case err != nil:
		return res, fmt.Errorf("unexpected database error: %w", err)
	}

	return res, nil
}
