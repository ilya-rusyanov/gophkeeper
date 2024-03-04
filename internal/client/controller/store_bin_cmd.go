package controller

import (
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// StoreBinCmd stores binary files
type StoreBinCmd struct {
	Name     string      `arg:"" help:"Name of object to be stored"`
	Meta     entity.Meta `help:"Meta information"`
	FilePath string      `arg:"" help:"Path to file"`
	uc       BinStorer
}

// Run executes the command
func (s *StoreBinCmd) Run(arg *Arg) error {
	err := s.uc.StoreBin(
		arg.Context,
		*entity.NewBinRecord(
			s.Name,
			s.Meta,
			[]byte{},
		),
		s.FilePath,
	)
	if err != nil {
		return fmt.Errorf("store bin use case failure: %s", err.Error())
	}

	return nil
}
