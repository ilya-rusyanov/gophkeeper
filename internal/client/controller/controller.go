package controller

import (
	"context"
)

// Registerer is registration domain
type Registerer interface {
	Register(login string, password string) error
}

// Controller is user controller for interaction with the application
type Controller struct {
	register Registerer
}

// Opt is a funcopt
type Opt func(*Controller)

func WithRegister(r Registerer) Opt {
	return func(c *Controller) {
		c.register = r
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

func (c *Controller) Run(ctx context.Context) error {
	// TODO: implement Run logic

	return nil
}
