package controller

import "fmt"

// ListCmd lists stored user data
type ListCmd struct {
	uc Lister
}

// Run executes the command
func (c *ListCmd) Run(arg *Arg) error {
	err := c.uc.List(arg.Context)
	if err != nil {
		return fmt.Errorf("usecase list error: %w", err)
	}

	return nil
}
