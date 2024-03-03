package list

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/list/mock"
)

//go:generate mockgen -destination ./mock/mocks.go -package mock . Servicer,Storager

func TestList(t *testing.T) {
	ctx := context.Background()

	t.Run("listing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		gw := mock.NewMockServicer(ctrl)
		storage := mock.NewMockStorager(ctrl)

		storage.EXPECT().
			Load().
			Return(entity.NewMyAuthentication("my auth"), nil)

		gw.EXPECT().
			List(gomock.Any(), entity.NewMyAuthentication("my auth")).
			Return(entity.DataList{
				entity.NewDataListEntry("auth", "yandex"),
				entity.NewDataListEntry("card", "tinkoff"),
			}, nil)

		uc := New(gw, storage)

		entries, err := uc.List(ctx)

		require.NoError(t, err)
		assert.Equal(t, entity.DataList{
			entity.NewDataListEntry("auth", "yandex"),
			entity.NewDataListEntry("card", "tinkoff"),
		}, entries)
	})
}
