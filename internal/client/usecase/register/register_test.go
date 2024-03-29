package register

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register/mock"
)

//go:generate mockgen -package mock -destination ./mock/servicer.go . Servicer
//go:generate mockgen -package mock -destination ./mock/storager.go . Storager

func TestRegister(t *testing.T) {
	user := "user"
	password := "password"
	someErr := errors.New("generic error")

	ctx := context.Background()

	t.Run("registration ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mock.NewMockServicer(ctrl)
		mockStorage := mock.NewMockStorager(ctrl)

		mockService.EXPECT().
			Register(gomock.Any(),
				*entity.NewMyCredentials(user, password)).
			Return(entity.NewMyAuthentication("auth"), nil)
		mockStorage.EXPECT().
			Store(gomock.Any(),
				entity.NewMyAuthentication("auth"))

		reg := New(mockService, mockStorage)

		err := reg.Register(ctx, *entity.NewMyCredentials(user, password))
		assert.NoError(t, err)
	})

	t.Run("registration failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mock.NewMockServicer(ctrl)
		mockStorage := mock.NewMockStorager(ctrl)

		mockService.EXPECT().
			Register(gomock.Any(),
				*entity.NewMyCredentials(user, password)).
			Return(entity.NewMyAuthentication(""), someErr)

		reg := New(mockService, mockStorage)

		err := reg.Register(ctx, *entity.NewMyCredentials(user, password))

		var ge *GenericError
		assert.ErrorAs(t, err, &ge)
	})

	t.Run("storage failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mock.NewMockServicer(ctrl)
		mockStorage := mock.NewMockStorager(ctrl)

		mockService.EXPECT().
			Register(gomock.Any(),
				*entity.NewMyCredentials(user, password))

		mockStorage.EXPECT().
			Store(gomock.Any(), gomock.Any()).
			Return(errors.New("storage faulure"))

		reg := New(mockService, mockStorage)

		err := reg.Register(ctx, *entity.NewMyCredentials(user, password))

		var ge *GenericError
		assert.ErrorAs(t, err, &ge)
	})
}
