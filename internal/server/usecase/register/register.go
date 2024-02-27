package register

import (
	"context"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
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
	salt string
	repo Repository
	log  Logger
}

// New constructs registration use case
func New(salt string, repo Repository, log Logger) *UC {
	return &UC{
		salt: salt,
		repo: repo,
		log:  log,
	}
}

// Register performs user registration
func (uc *UC) Register(
	ctx context.Context, creds entity.UserCredentials,
) error {
	// hash password with given salt
	dk, err := scrypt.Key([]byte(creds.Password), []byte(uc.salt), 32768, 8, 1, 32)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	creds.Password = hex.EncodeToString(dk)

	err = uc.repo.Store(ctx, creds)
	if err != nil {
		return fmt.Errorf("repository failed to store credentials: %w", err)
	}
	uc.log.Debugf("registered new user %q", creds.Login)

	return nil
}
