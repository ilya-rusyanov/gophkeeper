package controller

import (
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// StoreAuthCmd is a command for storing arbitrary auth data
type StoreAuthCmd struct {
	Name     string      `arg:"" help:"Name for object"`
	Meta     entity.Meta `help:"Meta information"`
	Login    string      `short:"l" help:"Login"`
	Password string      `short:"p" help:"Password"`
	uc       Storer      `kong:"-"`
}

func (s *StoreAuthCmd) Run(arg *Arg) error {
	err := s.uc.Store(
		arg.Context,
		*entity.NewAuthRecord(
			s.Name,
			s.Meta,
			entity.NewAuthPayload(
				s.Login,
				s.Password,
			),
		),
	)
	if err != nil {
		return fmt.Errorf("store auth use case failure: %s", err.Error())
	}

	return nil
}
