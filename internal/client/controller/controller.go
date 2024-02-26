package controller

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"
)

// Arg is execution context
type Arg struct {
	Context context.Context
}

// Registerer is registration domain
type Registerer interface {
	Register(ctx context.Context, login string, password string) error
}

// Controller is user controller for interaction with the application
type Controller struct {
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

	return &res
}

// Run starts the controller
func (c *Controller) Run(ctx context.Context) error {
	kongCtx := kong.Parse(&c.cli)

	if err := kongCtx.Run(&Arg{Context: ctx}); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	return nil
}
