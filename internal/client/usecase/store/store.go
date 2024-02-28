package store

import (
	"context"
	"errors"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// UC is data store use case
type UC struct{}

// New constructs new data store use case
func New() *UC {
	return &UC{}
}

// StoreAuth stores authentication data
func (uc *UC) StoreAuth(
	ctx context.Context, creds entity.Credentials,
) error {
	return errors.New("TODO")
}
