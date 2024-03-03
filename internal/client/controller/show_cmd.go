package controller

import (
	"fmt"
	"io"
	"strings"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

// ShowCmd is a command for revealing data
type ShowCmd struct {
	Type   string `arg:"" enum:"auth,card,text,bin" help:"Type of data"`
	Name   string `arg:"" help:"Name of stored data"`
	uc     Shower
	output io.Writer
}

// Run executes the command
func (c *ShowCmd) Run(arg *Arg) error {
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

	switch v := s.Payload.(type) {
	case entity.AuthPayload:
		buf.WriteString(`
login:		` + v.Login + `
password:	` + v.Password + "\n")
	default:
		return fmt.Errorf("unknown data type")
	}

	fmt.Fprint(c.output, buf.String())

	return nil
}
