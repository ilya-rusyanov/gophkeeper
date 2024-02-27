package register

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/register/mock"
)

//go:generate mockgen -package mock -destination ./mock/mocks.go . Repository

func TestRegister(t *testing.T) {
	t.Run("register ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockRepository(ctrl)

		repo.EXPECT().Store(gomock.Any(), entity.UserCredentials{
			Login:    "john",
			Password: "b8bad5db5f36d0fcd702445eb4d0c6b9f013c38035bba4cef62da2f2cb18b1f9",
		})

		reg := New("salt", repo)

		reg.Register(context.Background(), *entity.NewUserCredentials("john", "strongpw"))
	})
}
