package login

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/login/mock"
)

//go:generate mockgen -destination ./mock/storager.go -package mock . Storager
//go:generate mockgen -destination ./mock/tokener.go -package mock . Tokener

func TestUC(t *testing.T) {
	ctx := context.Background()

	t.Run("login ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storage := mock.NewMockStorager(ctrl)
		tokenBuilder := mock.NewMockTokener(ctrl)

		storage.EXPECT().
			GetByUsername(gomock.Any(), "john").
			Return("b8bad5db5f36d0fcd702445eb4d0c6b9f013c38035bba4cef62da2f2cb18b1f9", nil)

		tokenBuilder.EXPECT().
			Build(time.Second, "john").
			Return(entity.AuthToken("auth"), nil)

		uc := New(storage, "salt", tokenBuilder, time.Second)

		token, err := uc.LogIn(ctx, *entity.NewUserCredentials("john", "strongpw"))

		require.NoError(t, err)
		assert.Equal(t, entity.AuthToken("auth"), token)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storage := mock.NewMockStorager(ctrl)
		tokenBuilder := mock.NewMockTokener(ctrl)

		storage.EXPECT().
			GetByUsername(gomock.Any(), "john").
			Return("", entity.ErrNoSuchUser)

		uc := New(storage, "salt", tokenBuilder, time.Second)

		_, err := uc.LogIn(
			ctx,
			*entity.NewUserCredentials("john", "strongpw"),
		)

		assert.ErrorIs(t, err, entity.ErrNoSuchUser)
	})

	t.Run("wrong password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storage := mock.NewMockStorager(ctrl)
		tokenBuilder := mock.NewMockTokener(ctrl)

		storage.EXPECT().
			GetByUsername(gomock.Any(), "john").
			Return("this hash is different", nil)

		uc := New(storage, "salt", tokenBuilder, time.Second)

		_, err := uc.LogIn(
			ctx,
			*entity.NewUserCredentials("john", "strongpw"),
		)

		assert.ErrorIs(t, err, entity.ErrWrongPassword)
	})
}
