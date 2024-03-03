package show

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// Storager is user's authentication data storage
type Storager interface {
	Load() (entity.MyAuthentication, error)
}

// Servicer is remote service gateway
type Servicer interface {
	Show(context.Context, entity.ServiceShowRequest) (entity.Record, error)
}

// UC is use case for revealing data
type UC struct {
	storage Storager
	service Servicer
}

// New constructs the use case
func New(storage Storager, service Servicer) *UC {
	return &UC{
		storage: storage,
		service: service,
	}
}

// Show reveals user's data
func (uc *UC) Show(ctx context.Context, in entity.ShowIn) (entity.Record, error) {
	var res entity.Record

	auth, err := uc.storage.Load()
	if err != nil {
		return res, fmt.Errorf("failed to load auth data: %w", err)
	}

	res, err = uc.service.Show(ctx, entity.ServiceShowRequest{
		AuthData: auth,
		Type:     in.Type,
		Name:     in.Name,
	})
	if err != nil {
		return res, fmt.Errorf("gateway failure: %w", err)
	}

	return res, nil
}
