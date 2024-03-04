package controller

import (
	"context"
	"fmt"
	"io"

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

// BinStorer allows to store binary data
type BinStorer interface {
	StoreBin(ctx context.Context, rec entity.Record, filePath string) error
}

// LogIner is domain logic of loging user in
type LogIner interface {
	LogIn(context.Context, entity.MyCredentials) error
}

// Lister is stored data listing use case
type Lister interface {
	List(context.Context) (entity.DataList, error)
}

// Shower is usecase for showing stored information
type Shower interface {
	Show(context.Context, entity.ShowIn) (entity.Record, error)
	ShowBin(context.Context, entity.ShowBinIn) error
}

// Controller is user controller for interaction with the application
type Controller struct {
	cli struct {
		Register RegisterCmd `cmd:"" help:"Register user"`
		LogIn    LogInCmd    `cmd:"" help:"Log user in"`
		Store    StoreCmd    `cmd:"" help:"Store data"`
		List     ListCmd     `cmd:"" help:"List stored data"`
		Show     ShowCmd     `cmd:"" help:"Show stored data"`
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
func WithStore(s Storer) Opt {
	return func(c *Controller) {
		c.cli.Store.Auth.uc = s
		c.cli.Store.Text.uc = s
	}
}

// WithBinStore specifies binary Store use case
func WithBinStore(s BinStorer) Opt {
	return func(c *Controller) {
		c.cli.Store.Bin.uc = s
	}
}

// WithLogin supplies log in use case
func WithLogIn(l LogIner) Opt {
	return func(c *Controller) {
		c.cli.LogIn.uc = l
	}
}

// WitList supplies list use case
func WithList(l Lister) Opt {
	return func(c *Controller) {
		c.cli.List.uc = l
	}
}

// WithOutput supplies controller text output
func WithOutput(o io.Writer) Opt {
	return func(c *Controller) {
		c.cli.List.output = o
		c.cli.Show.output = o
	}
}

// WithShow specifies show use case
func WithShow(s Shower) Opt {
	return func(c *Controller) {
		c.cli.Show.uc = s
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
