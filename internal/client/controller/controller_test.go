package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/controller/mock"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

//go:generate mockgen -destination ./mock/registerer.go -package mock . Registerer

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

var someServerAddr = "localhost:8087"

func TestController(t *testing.T) {
	ctx := context.Background()

	t.Run("register", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reg := mock.NewMockRegisterer(ctrl)

		reg.EXPECT().Register(
			gomock.Any(),
			*entity.NewMyCredentials("john", "strongpw"),
		)

		c := New(args(t,
			"register",
			"--username", "john",
			"--password", "strongpw",
		))

		err := c.Run(ctx, WithRegister(reg))

		assert.NoError(t, err)
	})
}

func args(t testing.TB, a ...string) []string {
	res := []string{"--remote", someServerAddr}

	t.Helper()

	res = append(res, a...)

	return res
}
