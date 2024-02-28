package controller

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

type AuthStorer interface {
	StoreAuth(context.Context, entity.Credentials) error
}

// StoreAuthCmd is a command for storing arbitrary auth data
type StoreAuthCmd struct {
	Name     string     `arg:"" help:"Name for object"`
	Meta     []string   `help:"Meta information"`
	Login    string     `short:"l" help:"Login"`
	Password string     `short:"p" help:"Password"`
	uc       AuthStorer `kong:"-"`
}

func (s *StoreAuthCmd) Run(arg *Arg) error {
	err := s.uc.StoreAuth(arg.Context, entity.Credentials{
		Name:     s.Name,
		Meta:     s.Meta,
		Login:    s.Login,
		Password: s.Password,
	})
	if err != nil {
		return fmt.Errorf("store auth use case failure: %s", err.Error())
	}

	return nil
}
