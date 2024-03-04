package controller

import (
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// StoreTextCmd is a command for storing arbitrary auth data
type StoreTextCmd struct {
	Name string      `arg:"" help:"Name for object"`
	Meta entity.Meta `help:"Meta information"`
	Text string      `arg:"" help:"Arbitrary text"`
	uc   Storer
}

// Run executes the command
func (s *StoreTextCmd) Run(arg *Arg) error {
	err := s.uc.Store(
		arg.Context,
		*entity.NewTextRecord(
			s.Name,
			s.Meta,
			entity.NewTextPayload(
				s.Text,
			),
		),
	)
	if err != nil {
		return fmt.Errorf("store text use case failure: %s", err.Error())
	}

	return nil
}
