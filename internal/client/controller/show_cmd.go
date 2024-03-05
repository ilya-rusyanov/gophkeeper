package controller

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// ShowCmd is a command for revealing data
type ShowCmd struct {
	Type   string `arg:"" enum:"auth,card,text,bin" help:"Type of data"`
	Name   string `arg:"" help:"Name of stored data"`
	SaveTo string `name:"save-to" help:"For binary: path where to save data"`
	uc     Shower
	output io.Writer
}

// Run executes the command
func (c *ShowCmd) Run(arg *Arg) error {
	if c.Type == "bin" {
		return c.RunBin(arg)
	}

	if len(c.SaveTo) > 0 {
		return fmt.Errorf("--save-to is only for binary data")
	}

	s, err := c.uc.Show(arg.Context, entity.ShowIn{
		Type: entity.RecordType(c.Type),
		Name: c.Name,
	})
	if err != nil {
		return fmt.Errorf("usecase failed: %w", err)
	}

	buf := strings.Builder{}
	buf.WriteString(`meta:		`)

	for i, m := range s.Meta {
		buf.WriteString(`"` + m + `"`)

		if i != len(s.Meta)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteRune('\n')

	switch v := s.Payload.(type) {
	case entity.AuthPayload:
		buf.WriteString(`login:		` + v.Login + `
password:	` + v.Password + "\n")
	case entity.TextPayload:
		buf.WriteString(`text:		`)
		buf.WriteString(string(v))
		buf.WriteString("\n")
	case entity.CardPayload:
		buf.WriteString("number:\t\t" + v.Number + "\n")
		buf.WriteString("expires:\t" + strconv.Itoa(v.ExpMonth) + "/" + strconv.Itoa(v.ExpYear) + "\n")
		buf.WriteString("holder:\t\t" + v.HolderName + "\n")
		buf.WriteString("cvc:\t\t" + strconv.Itoa(v.CVC) + "\n")
	default:
		return fmt.Errorf("unknown data type to show")
	}

	fmt.Fprint(c.output, buf.String())

	return nil
}

func (c *ShowCmd) RunBin(arg *Arg) error {
	if len(c.SaveTo) == 0 {
		return fmt.Errorf("you have to specify --save-to")
	}

	err := c.uc.ShowBin(arg.Context,
		entity.ShowBinIn{
			Name:   c.Name,
			SaveTo: c.SaveTo,
		},
	)
	if err != nil {
		return fmt.Errorf("usecase failed to show bin: %w", err)
	}

	return nil
}
