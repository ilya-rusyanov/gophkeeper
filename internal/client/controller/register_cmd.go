package controller

import (
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// RegisterCmd is registration subcommand
type RegisterCmd struct {
	Username string     `required:"" help:"username"`
	Password string     `required:"" help:"password"`
	uc       Registerer `kong:"-"`
}

// Run performs registration
func (r *RegisterCmd) Run(arg *Arg) error {
	err := r.uc.Register(
		arg.Context,
		*entity.NewMyCredentials(r.Username, r.Password),
	)
	if err != nil {
		return fmt.Errorf("usecase registration error: %w", err)
	}

	return nil
}
