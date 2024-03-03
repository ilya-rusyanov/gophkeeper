package list

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// Servicer is remote gophkeeper service
type Servicer interface {
	List(context.Context, entity.MyAuthentication) (entity.DataList, error)
}

// Storager is authentication storage
type Storager interface {
	Load() (entity.MyAuthentication, error)
}

// UC is data listing use case
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

func (uc *UC) List(ctx context.Context) (entity.DataList, error) {
	var res entity.DataList

	myAuth, err := uc.storage.Load()
	if err != nil {
		return res, fmt.Errorf("auth storage failed to load auth: %w", err)
	}

	res, err = uc.service.List(ctx, myAuth)
	if err != nil {
		return res, fmt.Errorf("remote service failed to list data: %w", err)
	}

	return res, nil
}
