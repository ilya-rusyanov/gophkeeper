package controller

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"
)

// Registerer is registration domain
type Registerer interface {
	Register(login string, password string) error
}

// Controller is user controller for interaction with the application
type Controller struct {
	ctx *kong.Context
	cli struct {
		Register RegisterCmd `cmd:"" help:"Register user"`
	}
}

// Opt is a funcopt
type Opt func(*Controller)

// WithRegister specifies Register UC
func WithRegister(r Registerer) Opt {
	return func(c *Controller) {
		c.cli.Register.uc = r
	}
}

// New constructs controller
func New(opts ...Opt) *Controller {
	var res Controller

	for _, o := range opts {
		o(&res)
	}

	res.ctx = kong.Parse(&res.cli)

	return &res
}

// Run starts the controller
func (c *Controller) Run(ctx context.Context) error {
	if err := c.ctx.Run(); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	return nil
}
