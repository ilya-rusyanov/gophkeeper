package controller

import (
	"fmt"
)

// RegisterCmd is registration commnand
type RegisterCmd struct {
	Username string     `required:"" help:"username"`
	Password string     `required:"" help:"password"`
	uc       Registerer `kong:"-"`
}

// Run performs registration
func (r *RegisterCmd) Run(ctx *Arg) error {
	err := r.uc.Register(ctx.Context, r.Username, r.Password)
	if err != nil {
		return fmt.Errorf("usecase registration error: %w", err)
	}

	return nil
}
