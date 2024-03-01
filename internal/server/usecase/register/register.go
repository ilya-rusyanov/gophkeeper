package register

import (
	"context"
	"fmt"
	"time"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/pwhash"
)

type Logger interface {
	Debugf(string, ...any)
}

// Tokener builds authentication tokens
type Tokener interface {
	Build(lifespan time.Duration, login string) (entity.AuthToken, error)
}

// Repository stores user credentials
type Repository interface {
	Store(context.Context, entity.UserCredentials) error
}

// UC is registration use case
type UC struct {
	passwordSalt         string
	repo                 Repository
	log                  Logger
	tokenDefaultLifespan time.Duration
	tokenBuilder         Tokener
}

// New constructs registration use case
func New(
	passwordSalt string,
	repo Repository,
	log Logger,
	tokenDefaultLifespan time.Duration,
	tokenBuilder Tokener,
) *UC {
	return &UC{
		passwordSalt:         passwordSalt,
		repo:                 repo,
		log:                  log,
		tokenDefaultLifespan: tokenDefaultLifespan,
		tokenBuilder:         tokenBuilder,
	}
}

// Register performs user registration
func (uc *UC) Register(
	ctx context.Context, creds entity.UserCredentials,
) (entity.AuthToken, error) {
	var (
		authToken entity.AuthToken
		err       error
	)

	creds.Password, err = pwhash.Hash(uc.passwordSalt, creds.Password)
	if err != nil {
		return authToken, fmt.Errorf("failed to hash password: %w", err)
	}

	err = uc.repo.Store(ctx, creds)
	if err != nil {
		return authToken, fmt.Errorf("repository failed to store credentials: %w", err)
	}
	uc.log.Debugf("registered new user %q", creds.Login)

	authToken, err = uc.tokenBuilder.Build(
		uc.tokenDefaultLifespan,
		creds.Login,
	)
	if err != nil {
		return authToken, fmt.Errorf("failed to build auth token: %w", err)
	}

	return authToken, nil
}
