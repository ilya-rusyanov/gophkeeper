package register

import (
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/register/mock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -package mock -destination ./mock/repository.go . Repository
//go:generate mockgen -package mock -destination ./mock/tokener.go . Tokener

type dummyLog struct{}

func (l *dummyLog) Debugf(string, ...any) {}

func TestRegister(t *testing.T) {
	t.Run("register ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockRepository(ctrl)
		tokenBuilder := mock.NewMockTokener(ctrl)

		repo.EXPECT().Store(
			gomock.Any(),
			*entity.NewUserCredentials(
				"john",
				"b8bad5db5f36d0fcd702445eb4d0c6b9f013c38035bba4cef62da2f2cb18b1f9",
			))

		tokenBuilder.EXPECT().
			Build(time.Second, "john").
			Return(entity.AuthToken("auth"), nil)

		reg := New(
			"salt",
			repo,
			&dummyLog{},
			time.Second,
			tokenBuilder,
		)

		_, err := reg.Register(context.Background(), *entity.NewUserCredentials("john", "strongpw"))

		assert.NoError(t, err)
	})

	t.Run("user already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockRepository(ctrl)
		tokenBuilder := mock.NewMockTokener(ctrl)

		repo.EXPECT().Store(
			gomock.Any(),
			*entity.NewUserCredentials(
				"john",
				"b8bad5db5f36d0fcd702445eb4d0c6b9f013c38035bba4cef62da2f2cb18b1f9",
			)).Return(entity.ErrUserAlreadyExists)

		reg := New("salt", repo, &dummyLog{}, time.Second, tokenBuilder)

		_, err := reg.Register(context.Background(), *entity.NewUserCredentials("john", "strongpw"))

		assert.ErrorIs(t, err, entity.ErrUserAlreadyExists)
	})
}
