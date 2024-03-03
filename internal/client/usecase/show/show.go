package show

import (
	"context"
	"errors"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// UC is use case for revealing data
type UC struct{}

// New constructs the use case
func New() *UC {
	return &UC{}
}

// Show reveals user's data
func (uc *UC) Show(ctx context.Context, in entity.ShowIn) (entity.Record, error) {
	var res entity.Record

	return res, errors.New("TODO")
}
