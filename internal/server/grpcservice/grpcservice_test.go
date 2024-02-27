package grpcservice

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/grpcservice/mock"
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

//go:generate mockgen -destination ./mock/registeruc.go -package mock . RegisterUC

type dummyLogger struct{}

func (l *dummyLogger) Debugf(string, ...any) {}

var _ Logger = (*dummyLogger)(nil)

func TestRegister(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockRegisterUC(ctrl)

		uc.EXPECT().Register(
			gomock.Any(),
			*entity.NewUserCredentials("john", "strongpw"),
		)

		grpcsvc := New(&dummyLogger{}, uc)

		_, err := grpcsvc.Register(ctx, &proto.UserCredentials{
			Login:    "john",
			Password: "strongpw",
		})

		assert.NoError(t, err)
	})

	t.Run("user already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockRegisterUC(ctrl)

		uc.EXPECT().Register(gomock.Any(), gomock.Any()).
			Return(entity.ErrUserAlreadyExists)

		grpcsvc := New(&dummyLogger{}, uc)

		_, err := grpcsvc.Register(ctx, &proto.UserCredentials{})

		assert.Equal(t, codes.AlreadyExists, status.Code(err))
	})
}
