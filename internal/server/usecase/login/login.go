package login

import (
	"context"
	"fmt"
	"time"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/pwhash"
)

// Storage is storage for user's credential data
type Storager interface {
	GetByUsername(ctx context.Context, username string) (password string, err error)
}

// Tokener builds authentication tokens
type Tokener interface {
	Build(lifespan time.Duration, login string) (entity.AuthToken, error)
}

// UC is use case of user log in
type UC struct {
	storage              Storager
	passwordSalt         string
	tokenBuilder         Tokener
	tokenDefaultLifespan time.Duration
}

// New constructs the use case
func New(
	storage Storager,
	passwordSalt string,
	tokenBuilder Tokener,
	tokenLifetime time.Duration,
) *UC {
	return &UC{
		storage:              storage,
		passwordSalt:         passwordSalt,
		tokenBuilder:         tokenBuilder,
		tokenDefaultLifespan: tokenLifetime,
	}
}

// LogIn logs user in
func (uc *UC) LogIn(
	ctx context.Context, creds entity.UserCredentials,
) (entity.AuthToken, error) {
	var res entity.AuthToken

	hashedIn, err := pwhash.Hash(uc.passwordSalt, creds.Password)
	if err != nil {
		return res, fmt.Errorf("failed to hash password: %w", err)
	}

	storedHash, err := uc.storage.GetByUsername(ctx, creds.Login)
	if err != nil {
		return res, err
	}

	if hashedIn != storedHash {
		return res, entity.ErrWrongPassword
	}

	res, err = uc.tokenBuilder.Build(
		uc.tokenDefaultLifespan,
		creds.Login,
	)
	if err != nil {
		return res, fmt.Errorf("failed to build auth token: %w", err)
	}

	return res, nil
}
