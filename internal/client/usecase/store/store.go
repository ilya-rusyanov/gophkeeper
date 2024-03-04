package store

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// UC is data storing use case
type UC struct {
	authStorage AuthStorager
	service     Servicer
}

// New constructs the use case
func New(authStorage AuthStorager, service Servicer) *UC {
	return &UC{
		authStorage: authStorage,
		service:     service,
	}
}

// Store stores data in storage
func (uc *UC) Store(
	ctx context.Context, rec entity.Record,
) error {
	auth, err := uc.authStorage.Load()
	if err != nil {
		return fmt.Errorf("failed to load auth data: %w", err)
	}

	err = uc.service.Store(ctx,
		*entity.NewServiceStoreRequest(
			auth, rec,
		),
	)
	if err != nil {
		return fmt.Errorf("service failed to store data: %w", err)
	}

	return nil
}
