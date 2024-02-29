package register

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// GenericError represents generic error
type GenericError struct {
	orig error
}

// NewGenericError constructs GenericError
func NewGenericError(orig error) *GenericError {
	return &GenericError{
		orig: orig,
	}
}

// Error returns error text
func (e *GenericError) Error() string {
	return "generic error: " + e.orig.Error()
}

// Servicer is remote service
type Servicer interface {
	Register(context.Context, entity.MyCredentials) (entity.MyAuthentication, error)
}

// Storager stores user's authentication
type Storager interface {
	Store(context.Context, entity.MyAuthentication) error
}

// UC is use case for user registration on server
type UC struct {
	service Servicer
	storage Storager
}

// New constructs UC
func New(
	service Servicer,
	storage Storager,
) *UC {
	return &UC{
		service: service,
		storage: storage,
	}
}

// Register performs user registration
func (uc *UC) Register(ctx context.Context, credentials entity.MyCredentials) error {
	auth, err := uc.service.Register(ctx, credentials)
	if err != nil {
		return NewGenericError(
			fmt.Errorf("server error: %w", err),
		)
	}

	err = uc.storage.Store(ctx, auth)
	if err != nil {
		return NewGenericError(
			fmt.Errorf("failed to store my auth data: %w", err),
		)
	}

	return nil
}
