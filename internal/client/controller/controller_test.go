package controller

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/controller/mock"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

//go:generate mockgen -destination ./mock/mocks.go -package mock . Registerer,Storer,BinStorer,LogIner,Lister,Shower

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

	t.Run("log in", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		login := mock.NewMockLogIner(ctrl)

		login.EXPECT().LogIn(
			gomock.Any(),
			*entity.NewMyCredentials("john", "strongpw"),
		)

		c := New(args(t,
			"log-in",
			"--username", "john",
			"--password", "strongpw",
		))

		err := c.Run(ctx, WithLogIn(login))

		assert.NoError(t, err)
	})

	t.Run("store auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mock.NewMockStorer(ctrl)

		stor.EXPECT().Store(
			gomock.Any(),
			*entity.NewAuthRecord(
				"yandex mail",
				entity.Meta{
					"website:mail.ya.ru",
					"expires:june",
				},
				entity.NewAuthPayload("elon", "twitterx"),
			),
		)

		c := New(args(t,
			"store",
			"auth",
			"yandex mail",
			"--meta", "website:mail.ya.ru",
			"--meta", "expires:june",
			"-l", "elon",
			"-p", "twitterx",
		))

		err := c.Run(ctx, WithStore(stor))

		assert.NoError(t, err)
	})

	t.Run("store text", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mock.NewMockStorer(ctrl)

		stor.EXPECT().Store(
			gomock.Any(),
			*entity.NewTextRecord(
				"fish",
				entity.Meta{
					"topic:lost",
				},
				entity.NewTextPayload("Lorem ipsum dolor sit amet"),
			),
		)

		c := New(args(t,
			"store",
			"text",
			"fish",
			"--meta", "topic:lost",
			"Lorem ipsum dolor sit amet",
		))

		err := c.Run(ctx, WithStore(stor))

		assert.NoError(t, err)
	})

	t.Run("store bin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mock.NewMockBinStorer(ctrl)

		stor.EXPECT().StoreBin(
			gomock.Any(),
			*entity.NewBinRecord(
				"img",
				entity.Meta{
					"theme:sea",
				},
				[]byte{},
			),
			"/tmp/view.jpeg",
		)

		c := New(args(t,
			"store",
			"bin",
			"img",
			"--meta", "theme:sea",
			"/tmp/view.jpeg",
		))

		err := c.Run(ctx, WithBinStore(stor))

		assert.NoError(t, err)
	})

	t.Run("store card", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mock.NewMockStorer(ctrl)

		stor.EXPECT().Store(
			gomock.Any(),
			*entity.NewCardRecord(
				"tinkoff",
				entity.Meta{
					"color:black",
				},
				entity.NewCardPayload(
					"4377838623715638",
					6, 29,
					"MAGNUS CARLSEN",
					737),
			),
		)

		c := New(args(t,
			"store",
			"card",
			"tinkoff",
			"--meta", "color:black",
			"-n", "4377838623715638",
			"-e", "06/29",
			"-o", "MAGNUS CARLSEN",
			"-c", "737",
		))

		err := c.Run(ctx, WithStore(stor))

		assert.NoError(t, err)
	})

	t.Run("listing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		list := mock.NewMockLister(ctrl)

		output := strings.Builder{}

		list.EXPECT().List(
			gomock.Any(),
		).Return(
			entity.DataList{
				entity.NewDataListEntry("auth", "yandex"),
				entity.NewDataListEntry("card", "tinkoff"),
			},
			nil,
		)

		c := New(args(t,
			"list",
		))

		err := c.Run(ctx,
			WithList(list),
			WithOutput(&output),
		)

		require.NoError(t, err)
		assert.Equal(t, `auth	"yandex"
card	"tinkoff"
`, output.String())
	})

	t.Run("show auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		show := mock.NewMockShower(ctrl)

		output := strings.Builder{}

		show.EXPECT().Show(
			gomock.Any(),
			entity.ShowIn{
				Type: "auth",
				Name: "yandex",
			},
		).
			Return(
				entity.Record{
					Type: entity.RecordTypeAuth,
					Name: "yandex",
					Meta: entity.Meta{
						"expire:july",
						"use:never",
					},
					Payload: entity.NewAuthPayload("elon", "twitterx"),
				},
				nil,
			)

		c := New(args(
			t,
			"show",
			"auth",
			"yandex",
		))

		err := c.Run(ctx,
			WithShow(show),
			WithOutput(&output),
		)

		require.NoError(t, err)
		assert.Equal(t, `meta:		"expire:july", "use:never"
login:		elon
password:	twitterx
`, output.String())
	})

	t.Run("show text", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		show := mock.NewMockShower(ctrl)

		output := strings.Builder{}

		show.EXPECT().Show(
			gomock.Any(),
			entity.ShowIn{
				Type: "text",
				Name: "fish",
			},
		).
			Return(
				entity.Record{
					Type: entity.RecordTypeText,
					Name: "fish",
					Meta: entity.Meta{
						"meaning:no",
					},
					Payload: entity.NewTextPayload("Lorem ipsum dolor sit amet"),
				},
				nil,
			)

		c := New(args(
			t,
			"show",
			"text",
			"fish",
		))

		err := c.Run(ctx,
			WithShow(show),
			WithOutput(&output),
		)

		require.NoError(t, err)
		assert.Equal(t, `meta:		"meaning:no"
text:		Lorem ipsum dolor sit amet
`, output.String())
	})

	t.Run("show binary", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		show := mock.NewMockShower(ctrl)

		show.EXPECT().
			ShowBin(
				gomock.Any(),
				entity.ShowBinIn{
					Name:   "img",
					SaveTo: "/tmp/view.jpeg",
				},
			).
			Return(nil)

		c := New(args(
			t,
			"show",
			"bin",
			"img",
			"--save-to", "/tmp/view.jpeg",
		))

		err := c.Run(ctx,
			WithShow(show),
		)

		require.NoError(t, err)
	})

	t.Run("show card", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		show := mock.NewMockShower(ctrl)

		output := strings.Builder{}

		show.EXPECT().Show(
			gomock.Any(),
			entity.ShowIn{
				Type: "card",
				Name: "tinkoff",
			},
		).
			Return(
				entity.Record{
					Type: entity.RecordTypeCard,
					Name: "tinkoff",
					Meta: entity.Meta{
						"color:black",
					},
					Payload: entity.NewCardPayload(
						"4377838623715638",
						6, 29,
						"MAGNUS CARLSEN",
						737,
					),
				},
				nil,
			)

		c := New(args(
			t,
			"show",
			"card",
			"tinkoff",
		))

		err := c.Run(ctx,
			WithShow(show),
			WithOutput(&output),
		)

		require.NoError(t, err)
		assert.Equal(t, `meta:		"color:black"
number:		4377838623715638
expires:	6/29
holder:		MAGNUS CARLSEN
cvc:		737
`, output.String())
	})
}

func args(t testing.TB, a ...string) []string {
	res := []string{"--remote", someServerAddr}

	t.Helper()

	res = append(res, a...)

	return res
}
