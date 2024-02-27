package user

import (
	"context"
	"errors"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
)

// Repo is repository for manipulating with users and their credentials
type Repo struct{}

func New() *Repo {
	return &Repo{}
}

// Store stores user credentials
func (r *Repo) Store(
	ctx context.Context, creds entity.UserCredentials,
) error {
	return errors.New("TODO")
}
