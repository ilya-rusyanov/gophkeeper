package register

import (
	"errors"
	"testing"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegister(t *testing.T) {
	user := "user"
	password := "password"
	someErr := errors.New("generic error")

	t.Run("registration ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCredStorage := mock.NewMockCredentialsStorager(ctrl)
		mockService := mock.NewMockServicer(ctrl)

		mockService.EXPECT().Register(user, password).Return(nil)
		mockCredStorage.EXPECT().Store(user, password).Return(nil)

		reg := New(mockCredStorage, mockService)

		err := reg.Register(user, password)
		assert.NoError(t, err)
	})

	t.Run("registration failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCredStorage := mock.NewMockCredentialsStorager(ctrl)
		mockService := mock.NewMockServicer(ctrl)

		mockService.EXPECT().Register(user, password).Return(someErr)

		reg := New(mockCredStorage, mockService)

		err := reg.Register(user, password)

		var ge *GenericError
		assert.ErrorAs(t, err, &ge)
	})

	t.Run("storage failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCredStorage := mock.NewMockCredentialsStorager(ctrl)
		mockService := mock.NewMockServicer(ctrl)

		mockService.EXPECT().Register(user, password).Return(nil)
		mockCredStorage.EXPECT().Store(user, password).Return(someErr)

		reg := New(mockCredStorage, mockService)

		err := reg.Register(user, password)

		var ge *GenericError
		assert.ErrorAs(t, err, &ge)
	})
}
