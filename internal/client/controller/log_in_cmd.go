package controller

import (
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// LogInCmd is logging in subcommand
type LogInCmd struct {
	Username string  `required:"" help:"username"`
	Password string  `required:"" help:"password"`
	uc       LogIner `kong:"-"`
}

// Run performs log in
func (r *LogInCmd) Run(arg *Arg) error {
	err := r.uc.LogIn(
		arg.Context,
		*entity.NewMyCredentials(r.Username, r.Password),
	)
	if err != nil {
		return fmt.Errorf("usecase log in error: %w", err)
	}

	return nil
}
