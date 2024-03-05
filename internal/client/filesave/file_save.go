package filesave

import (
	"context"
	"fmt"
	"os"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// FileSave is object for writing binary files
type FileSave struct{}

// New constructs FileSave object
func New() *FileSave {
	return &FileSave{}
}

// SaveFile writes data to disk
func (s *FileSave) SaveFile(_ context.Context, in entity.FileSaveIn) error {
	err := os.WriteFile(in.FilePath, []byte(*in.Data), 0o600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return err
}
