package controller

import (
	"fmt"
	"io"
)

// ListCmd lists stored user data
type ListCmd struct {
	uc     Lister
	output io.Writer
}

// Run executes the command
func (c *ListCmd) Run(arg *Arg) error {
	l, err := c.uc.List(arg.Context)
	if err != nil {
		return fmt.Errorf("usecase list error: %w", err)
	}

	for _, entry := range l {
		fmt.Fprintf(c.output, "%s\t%q\n", entry.Type, entry.Name)
	}

	return nil
}
