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

// Storer domain logic of storing data
type Storer interface {
	Store(context.Context, entity.Record) error
}

// Controller is user controller for interaction with the application
type Controller struct {
	cli struct {
		Register RegisterCmd `cmd:"" help:"Register user"`
		Store    StoreCmd    `cmd:"" help:"Store data"`
		Config   Config      `embed:""`
	}
	args []string
}

// Opt is a funcopt
type Opt func(*Controller)

// WithRegister specifies Register use case
func WithRegister(r Registerer) Opt {
	return func(c *Controller) {
		c.cli.Register.uc = r
	}
}

// WithStore specifies Store use case
func WithStore(r Storer) Opt {
	return func(c *Controller) {
		c.cli.Store.Auth.uc = r
	}
}

// ReadConfig reads and returns app configuration
func ReadConfig(args []string) Config {
	var ctrl Controller

	_ = parse(args, &ctrl.cli)

	return ctrl.cli.Config
}

// New constructs controller
func New(args []string) *Controller {
	var res Controller

	res.args = args

	return &res
}

// Run starts the controller
func (c *Controller) Run(
	ctx context.Context, opts ...Opt,
) error {
	kongCtx := parse(c.args, &c.cli)

	for _, o := range opts {
		o(c)
	}

	if err := kongCtx.Run(&Arg{Context: ctx}); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	return nil
}

func parse(args []string, grammar any) *kong.Context {
	k, err := kong.New(grammar)
	if err != nil {
		panic(err)
	}
	ctx, err := k.Parse(args)
	k.FatalIfErrorf(err)

	return ctx
}
