package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/store/mock"
)

//go:generate mockgen -destination ./mock/auth_storager.go -package mock . AuthStorager
//go:generate mockgen -destination ./mock/servicer.go -package mock . Servicer

func TestStore(t *testing.T) {
	ctx := context.Background()

	t.Run("store auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		authRepo := mock.NewMockAuthStorager(ctrl)
		service := mock.NewMockServicer(ctrl)

		authRepo.EXPECT().
			Load().
			Return(entity.NewMyAuthentication("auth"), nil)

		service.EXPECT().
			Store(gomock.Any(),
				*entity.NewServiceStoreRequest(
					entity.NewMyAuthentication(
						"auth",
					),
					*entity.NewAuthRecord(
						"yandex mail",
						entity.Meta{"expires:july"},
						entity.NewAuthPayload("john", "strongpw"),
					),
				),
			)

		uc := New(authRepo, service)

		err := uc.Store(ctx,
			*entity.NewAuthRecord(
				"yandex mail",
				entity.Meta{"expires:july"},
				entity.NewAuthPayload("john", "strongpw"),
			),
		)

		assert.NoError(t, err)
	})
}
