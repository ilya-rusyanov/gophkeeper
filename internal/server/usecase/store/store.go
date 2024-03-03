package store

import (
	"context"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
)

// Storer is storage for user's data
type Storer interface {
	Store(context.Context, *entity.StoreIn) error
}

// UC represents store data use case
type UC struct {
	storage Storer
}

// New constructs store data use case
func New(storage Storer) *UC {
	return &UC{
		storage: storage,
	}
}

func (uc *UC) Store(ctx context.Context, in *entity.StoreIn) error {
	return uc.storage.Store(ctx, in)
}
