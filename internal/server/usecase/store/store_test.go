package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/store/mock"
)

//go:generate mockgen -destination ./mock/storer.go -package mock . Storer

func TestUC(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull store", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storage := mock.NewMockStorer(ctrl)

		storage.EXPECT().
			Store(
				gomock.Any(),
				entity.NewStoreIn(
					"john",
					"auth",
					"yandex mail",
					`["expire:june"]`,
					[]byte(`{"login":"john", "password":"strongpw"}`),
				),
			)

		uc := New(storage)

		err := uc.Store(ctx,
			entity.NewStoreIn(
				"john",
				"auth",
				"yandex mail",
				`["expire:june"]`,
				[]byte(`{"login":"john", "password":"strongpw"}`),
			),
		)

		assert.NoError(t, err)
	})
}
