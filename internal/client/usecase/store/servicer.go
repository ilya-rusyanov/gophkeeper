package store

import (
	"context"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// Servicer is gophkeeper service gateway
type Servicer interface {
	Store(context.Context, entity.ServiceStoreRequest) error
}
