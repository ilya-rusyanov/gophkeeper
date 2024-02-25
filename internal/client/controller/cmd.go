package controller

import "fmt"

// RegisterCmd is registration commnand
type RegisterCmd struct {
	Username string     `required:"" help:"username"`
	Password string     `required:"" help:"password"`
	uc       Registerer `kong:"-"`
}

// Run runs registration command
func (r *RegisterCmd) Run() error {
	err := r.uc.Register(r.Username, r.Password)
	if err != nil {
		return fmt.Errorf("usecase registration error: %w", err)
	}

	return nil
}
