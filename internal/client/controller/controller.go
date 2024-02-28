package controller

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// Arg is execution context
type Arg struct {
	Context context.Context
}

// Registerer is registration domain logic
type Registerer interface {
	Register(context.Context, entity.MyCredentials) error
}

// Controller is user controller for interaction with the application
type Controller struct {
	cli struct {
		Register RegisterCmd `cmd:"" help:"Register user"`
		config   Config      `embed:""`
	}
	kongCtx *kong.Context
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
func New() (*Controller, Config) {
	var res Controller

	res.kongCtx = kong.Parse(&res.cli)

	return &res, res.cli.config
}

// Run starts the controller
func (c *Controller) Run(
	ctx context.Context, opts ...Opt,
) error {
	for _, o := range opts {
		o(c)
	}

	if err := c.kongCtx.Run(&Arg{Context: ctx}); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	return nil
}
