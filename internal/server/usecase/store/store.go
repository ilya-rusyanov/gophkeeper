package store

import (
	"context"
	"errors"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
)

// UC represents store data use case
type UC struct{}

// New constructs store data use case
func New() *UC {
	return &UC{}
}

func (uc *UC) Store(ctx context.Context, in *entity.StoreIn) error {
	return errors.New("TODO")
}
