package fileread

import (
	"context"
	"fmt"
	"os"
)

// FileRead is object that allows to read files
type FileRead struct{}

// New constructs FileRead
func New() *FileRead {
	return &FileRead{}
}

// ReadFile reads and returns file
func (r *FileRead) ReadFile(_ context.Context, filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return file, nil
}
