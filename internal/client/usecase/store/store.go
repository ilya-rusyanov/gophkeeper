package store

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// CredStorager is storage of user's own credentials
type CredStorager interface {
	Load() (entity.MyCredentials, error)
}

// Servicer is gophkeeper service gateway
type Servicer interface {
	Store(context.Context, entity.ServiceStoreRequest) error
}

// UC is data storing use case
type UC struct {
	credStorage CredStorager
	service     Servicer
}

// New constructs the use case
func New(credStore CredStorager, service Servicer) *UC {
	return &UC{
		credStorage: credStore,
		service:     service,
	}
}

// Store stores data in storage
func (uc *UC) Store(
	ctx context.Context, rec entity.Record,
) error {
	cred, err := uc.credStorage.Load()
	if err != nil {
		return fmt.Errorf("failed to load user credentials: %w", err)
	}

	err = uc.service.Store(ctx,
		*entity.NewServiceStoreRequest(
			cred, rec,
		),
	)
	if err != nil {
		return fmt.Errorf("service failed to store data: %w", err)
	}

	return nil
}
