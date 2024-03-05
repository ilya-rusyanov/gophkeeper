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

// FileSaver saves files to disk
type FileSaver interface {
	SaveFile(context.Context, entity.FileSaveIn) error
}

// UC is use case for revealing data
type UC struct {
	storage   Storager
	service   Servicer
	fileSaver FileSaver
}

// New constructs the use case
func New(storage Storager, service Servicer, fileSaver FileSaver) *UC {
	return &UC{
		storage:   storage,
		service:   service,
		fileSaver: fileSaver,
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

// ShowBin is for showing binary data
func (uc *UC) ShowBin(ctx context.Context, in entity.ShowBinIn) error {
	var rec entity.Record

	auth, err := uc.storage.Load()
	if err != nil {
		return fmt.Errorf("failed to load auth data: %w", err)
	}

	rec, err = uc.service.Show(ctx, entity.ServiceShowRequest{
		AuthData: auth,
		Type:     entity.RecordTypeBin,
		Name:     in.Name,
	})
	if err != nil {
		return fmt.Errorf("gateway failure: %w", err)
	}

	data, ok := rec.Payload.(*entity.BinPayload)
	if !ok {
		return fmt.Errorf("received incorrect data type")
	}

	err = uc.fileSaver.SaveFile(ctx, entity.FileSaveIn{
		Data:     data,
		FilePath: in.SaveTo,
	})
	if err != nil {
		return fmt.Errorf("failed to save file to disk: %w", err)
	}

	return nil
}
