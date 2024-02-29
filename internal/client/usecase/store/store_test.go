package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/store/mock"
)

//go:generate mockgen -destination ./mock/cred_storager.go -package mock . CredStorager
//go:generate mockgen -destination ./mock/servicer.go -package mock . Servicer

func TestStore(t *testing.T) {
	ctx := context.Background()

	t.Run("store auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		credRepo := mock.NewMockCredStorager(ctrl)
		service := mock.NewMockServicer(ctrl)

		credRepo.EXPECT().
			Load().
			Return(*entity.NewMyCredentials("john", "strongpw"), nil)

		service.EXPECT().
			Store(gomock.Any(),
				*entity.NewServiceStoreRequest(
					*entity.NewMyCredentials("john", "strongpw"),
					*entity.NewAuthRecord(
						"yandex mail",
						entity.Meta{"expires:july"},
						entity.NewAuthPayload("john", "strongpw"),
					),
				),
			)

		uc := New(credRepo, service)

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
