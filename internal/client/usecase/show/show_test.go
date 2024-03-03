package show

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/show/mock"
)

//go:generate mockgen -destination ./mock/mocks.go -package mock . Servicer,Storager

func TestShow(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull show", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storager := mock.NewMockStorager(ctrl)
		servicer := mock.NewMockServicer(ctrl)

		storager.EXPECT().
			Load().
			Return(entity.NewMyAuthentication("my auth"), nil)

		servicer.EXPECT().
			Show(gomock.Any(),
				entity.ServiceShowRequest{
					AuthData: entity.NewMyAuthentication("my auth"),
					Type:     entity.RecordTypeAuth,
					Name:     "yandex",
				}).Return(
			*entity.NewAuthRecord(
				"yandex",
				entity.Meta{"expires:july"},
				entity.NewAuthPayload("elon", "twitterx"),
			), nil)

		uc := New(storager, servicer)

		s, err := uc.Show(ctx, entity.ShowIn{
			Type: "auth",
			Name: "yandex",
		})

		require.NoError(t, err)
		assert.Equal(t, *entity.NewAuthRecord(
			"yandex",
			entity.Meta{"expires:july"},
			entity.NewAuthPayload("elon", "twitterx"),
		), s)
	})
}
