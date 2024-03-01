package register

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/pwhash"
)

type Logger interface {
	Debugf(string, ...any)
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
	tokenSigningKey      string
}

// New constructs registration use case
func New(
	passwordSalt string,
	repo Repository,
	log Logger,
	tokenDefaultLifespan time.Duration,
	tokenSigningKey string,
) *UC {
	return &UC{
		passwordSalt:         passwordSalt,
		repo:                 repo,
		log:                  log,
		tokenDefaultLifespan: tokenDefaultLifespan,
		tokenSigningKey:      tokenSigningKey,
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

	authToken, err = buildAuthToken(
		uc.tokenDefaultLifespan,
		creds.Login,
		uc.tokenSigningKey,
	)
	if err != nil {
		return authToken, fmt.Errorf("failed to build auth token: %w", err)
	}

	return authToken, nil
}

func buildAuthToken(
	expireIn time.Duration, login string, key string,
) (entity.AuthToken, error) {
	var result entity.AuthToken

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(expireIn)),
		},
		Login: login,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return result, fmt.Errorf("failed to sign token: %w", err)
	}

	result = entity.AuthToken(tokenString)

	return result, nil
}
