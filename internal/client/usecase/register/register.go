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
	Register(context.Context, entity.MyCredentials) error
}

// Register is UC for user registration on server
type Register struct {
	service Servicer
}

// New constructs UC
func New(
	service Servicer,
) *Register {
	return &Register{
		service: service,
	}
}

// Register performs user registration
func (r *Register) Register(ctx context.Context, credentials entity.MyCredentials) error {
	if err := r.service.Register(ctx, credentials); err != nil {
		return NewGenericError(
			fmt.Errorf("server error: %w", err),
		)
	}

	return nil
}
