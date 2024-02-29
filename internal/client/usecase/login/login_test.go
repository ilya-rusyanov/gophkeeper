package login

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/login/mock"
)

//go:generate mockgen -destination ./mock/servicer.go -package mock . Servicer
//go:generate mockgen -destination ./mock/storager.go -package mock . Storager

func TestUC(t *testing.T) {
	ctx := context.Background()

	t.Run("log in ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mock.NewMockServicer(ctrl)
		storage := mock.NewMockStorager(ctrl)

		service.EXPECT().
			LogIn(ctx, *entity.NewMyCredentials("john", "strongpw")).
			Return(entity.NewMyAuthentication("auth"), nil)

		storage.EXPECT().
			Store(ctx, entity.NewMyAuthentication("auth"))

		uc := New(service, storage)

		err := uc.LogIn(ctx, *entity.NewMyCredentials("john", "strongpw"))

		assert.NoError(t, err)
	})

	t.Run("log in failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mock.NewMockServicer(ctrl)
		storage := mock.NewMockStorager(ctrl)

		service.EXPECT().
			LogIn(ctx, *entity.NewMyCredentials("john", "strongpw")).
			Return(entity.NewMyAuthentication(""), errors.New("some error"))

		uc := New(service, storage)

		err := uc.LogIn(ctx, *entity.NewMyCredentials("john", "strongpw"))

		assert.NotNil(t, err)
	})

	t.Run("store failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := mock.NewMockServicer(ctrl)
		storage := mock.NewMockStorager(ctrl)

		service.EXPECT().
			LogIn(ctx, *entity.NewMyCredentials("john", "strongpw")).
			Return(entity.NewMyAuthentication(""), nil)

		storage.EXPECT().
			Store(ctx, entity.NewMyAuthentication("")).
			Return(errors.New("storage error"))

		uc := New(service, storage)

		err := uc.LogIn(ctx, *entity.NewMyCredentials("john", "strongpw"))

		assert.NotNil(t, err)
	})
}
