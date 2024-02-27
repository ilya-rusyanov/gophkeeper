package register

import (
	"context"
	"fmt"
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

// CredentialsStorager represents storage of credentials
type CredentialsStorager interface {
	Store(login, password string) error
}

// Servicer is remote service
type Servicer interface {
	Register(ctx context.Context, login, password string) error
}

// Register is UC for user registration on server
type Register struct {
	service Servicer
	storage CredentialsStorager
}

// New constructs UC
func New(
	credStorage CredentialsStorager,
	service Servicer,
) *Register {
	return &Register{
		service: service,
		storage: credStorage,
	}
}

// Register performs user registration
func (r *Register) Register(ctx context.Context, login, password string) error {
	if err := r.service.Register(ctx, login, password); err != nil {
		return NewGenericError(
			fmt.Errorf("server error: %w", err),
		)
	}

	if err := r.storage.Store(login, password); err != nil {
		return NewGenericError(
			fmt.Errorf("credentials storage error: %w", err),
		)
	}

	return nil
}
