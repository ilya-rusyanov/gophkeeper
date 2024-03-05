package controller

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

var ErrExpireSyntax = errors.New("expire must be in MONTH/YEAR form")

// StoreCardCmd is for storing card information
type StoreCardCmd struct {
	Name   string      `arg:"" help:"Name of card"`
	Meta   entity.Meta `help:"Meta information"`
	Number string      `required:"" short:"n" help:"Card number"`
	Expire string      `required:"" short:"e" help:"Card expiration MONTH/YEAR format"`
	Owner  string      `required:"" short:"o" help:"Card owner name"`
	CVC    int         `required:"" short:"c" help:"Card CVC number"`
	uc     Storer
}

// Run executes the command
func (c *StoreCardCmd) Run(arg *Arg) error {
	s := strings.Split(c.Expire, "/")
	if len(s) != 2 {
		return ErrExpireSyntax
	}

	month, err := strconv.Atoi(s[0])
	if err != nil {
		return ErrExpireSyntax
	}

	year, err := strconv.Atoi(s[1])
	if err != nil {
		return ErrExpireSyntax
	}

	err = c.uc.Store(
		arg.Context,
		*entity.NewCardRecord(
			c.Name,
			c.Meta,
			entity.NewCardPayload(
				c.Number,
				month,
				year,
				c.Owner,
				c.CVC,
			),
		),
	)
	if err != nil {
		return fmt.Errorf("store use case failure: %w", err)
	}

	return nil
}
