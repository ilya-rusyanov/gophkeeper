package register

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register/mock"
)

//go:generate mockgen -package mock -destination ./mock/mocks.go . Servicer,CredentialsStorager

func TestRegister(t *testing.T) {
	user := "user"
	password := "password"
	someErr := errors.New("generic error")

	ctx := context.Background()

	t.Run("registration ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCredStorage := mock.NewMockCredentialsStorager(ctrl)
		mockService := mock.NewMockServicer(ctrl)

		mockService.EXPECT().Register(gomock.Any(), user, password).Return(nil)
		mockCredStorage.EXPECT().Store(user, password).Return(nil)

		reg := New(mockCredStorage, mockService)

		err := reg.Register(ctx, user, password)
		assert.NoError(t, err)
	})

	t.Run("registration failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCredStorage := mock.NewMockCredentialsStorager(ctrl)
		mockService := mock.NewMockServicer(ctrl)

		mockService.EXPECT().Register(gomock.Any(), user, password).Return(someErr)

		reg := New(mockCredStorage, mockService)

		err := reg.Register(ctx, user, password)

		var ge *GenericError
		assert.ErrorAs(t, err, &ge)
	})

	t.Run("storage failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCredStorage := mock.NewMockCredentialsStorager(ctrl)
		mockService := mock.NewMockServicer(ctrl)

		mockService.EXPECT().Register(gomock.Any(), user, password).Return(nil)
		mockCredStorage.EXPECT().Store(user, password).Return(someErr)

		reg := New(mockCredStorage, mockService)

		err := reg.Register(ctx, user, password)

		var ge *GenericError
		assert.ErrorAs(t, err, &ge)
	})
}
