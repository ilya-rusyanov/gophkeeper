package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	t.Run("values", func(t *testing.T) {
		c := ReadConfig([]string{
			"--remote", "localhost:8089",
			"--verbosity", "error",
			"register",
			"--username", "a",
			"--password", "b",
		})

		assert.Equal(t, "localhost:8089", c.Server)
		assert.Equal(t, "error", c.LogLevel)
	})
}
