package store

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// FileReader allows to read files
type FileReader interface {
	ReadFile(ctx context.Context, fileName string) ([]byte, error)
}

// StoreBin is use case for storing binary data
type StoreBin struct {
	authStorage AuthStorager
	fileReader  FileReader
	service     Servicer
}

// NewBin constructs store binary use case
func NewBin(
	authStorage AuthStorager,
	fileReader FileReader,
	service Servicer,
) *StoreBin {
	return &StoreBin{
		authStorage: authStorage,
		fileReader:  fileReader,
		service:     service,
	}
}

// StoreBin stores binary file
func (s *StoreBin) StoreBin(
	ctx context.Context,
	recPrototype entity.Record,
	filePath string,
) error {
	auth, err := s.authStorage.Load()
	if err != nil {
		return fmt.Errorf("failed to load auth data: %w", err)
	}

	fileContents, err := s.fileReader.ReadFile(ctx, filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	payload := entity.BinPayload(fileContents)
	recPrototype.Payload = &payload

	err = s.service.Store(ctx,
		*entity.NewServiceStoreRequest(
			auth, recPrototype,
		),
	)
	if err != nil {
		return fmt.Errorf("service failed to store data: %w", err)
	}

	return nil
}
