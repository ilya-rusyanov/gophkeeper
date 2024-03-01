package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// Storage is storage for authentication data
type Storage struct {
	filename string
}

// New constructs storage
func New(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (s *Storage) Store(_ context.Context, data entity.MyAuthentication) error {
	if err := os.WriteFile(s.filename, []byte(data), 0o600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (s *Storage) Load() (entity.MyAuthentication, error) {
	var res entity.MyAuthentication

	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return res, entity.ErrAuthAbsent
		}

		return res, fmt.Errorf("unexpectd error reading authentication file: %w", err)
	}

	res = entity.NewMyAuthentication(string(data))

	return res, nil
}
