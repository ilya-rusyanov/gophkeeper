package login

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// Servicer is remote service
type Servicer interface {
	LogIn(context.Context, entity.MyCredentials) (entity.MyAuthentication, error)
}

// Storager is authentication data storage
type Storager interface {
	Store(context.Context, entity.MyAuthentication) error
}

// UC is log in use case
type UC struct {
	service Servicer
	storage Storager
}

// New constructs the use case
func New(service Servicer, storage Storager) *UC {
	return &UC{
		service: service,
		storage: storage,
	}
}

// LogIn logs user in
func (uc *UC) LogIn(ctx context.Context, cred entity.MyCredentials) error {
	auth, err := uc.service.LogIn(ctx, cred)
	if err != nil {
		return fmt.Errorf("remote service failed to log user in: %w", err)
	}

	err = uc.storage.Store(ctx, auth)
	if err != nil {
		return fmt.Errorf("failed to store auth data: %w", err)
	}

	return nil
}
