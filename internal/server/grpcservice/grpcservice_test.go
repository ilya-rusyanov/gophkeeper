package grpcservice

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/grpcservice/mock"
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

//go:generate mockgen -destination ./mock/iregisteruc.go -package mock . IRegisterUC
//go:generate mockgen -destination ./mock/loginer.go -package mock . LogIner
//go:generate mockgen -destination ./mock/istoreuc.go -package mock . IStoreUC

type dummyLogger struct{}

func (l *dummyLogger) Debugf(string, ...any) {}

var _ Logger = (*dummyLogger)(nil)

func TestRegister(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockIRegisterUC(ctrl)

		uc.EXPECT().Register(
			gomock.Any(),
			*entity.NewUserCredentials("john", "strongpw"),
		)

		grpcsvc := New(&dummyLogger{}, uc, nil, nil)

		_, err := grpcsvc.Register(ctx,
			&proto.RegisterRequest{
				Credentials: &proto.UserCredentials{
					Login:    "john",
					Password: "strongpw",
				},
			})

		assert.NoError(t, err)
	})

	t.Run("user already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockIRegisterUC(ctrl)

		uc.EXPECT().Register(gomock.Any(), gomock.Any()).
			Return(entity.AuthToken(""), entity.ErrUserAlreadyExists)

		grpcsvc := New(&dummyLogger{}, uc, nil, nil)

		_, err := grpcsvc.Register(
			ctx,
			&proto.RegisterRequest{
				Credentials: &proto.UserCredentials{},
			})

		assert.Equal(t, codes.AlreadyExists, status.Code(err))
	})

	t.Run("other error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockIRegisterUC(ctrl)

		uc.EXPECT().Register(gomock.Any(), gomock.Any()).
			Return(entity.AuthToken(""), errors.New("a different error"))

		grpcsvc := New(&dummyLogger{}, uc, nil, nil)

		_, err := grpcsvc.Register(ctx,
			&proto.RegisterRequest{
				Credentials: &proto.UserCredentials{},
			})

		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestLogIn(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull log in", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockLogIner(ctrl)

		uc.EXPECT().
			LogIn(gomock.Any(), *entity.NewUserCredentials("john", "strongpw"))

		grpcsvc := New(&dummyLogger{}, nil, uc, nil)

		_, err := grpcsvc.LogIn(ctx,
			&proto.LogInRequest{
				Credentials: &proto.UserCredentials{
					Login:    "john",
					Password: "strongpw",
				},
			},
		)

		assert.NoError(t, err)
	})

	t.Run("no such user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockLogIner(ctrl)

		uc.EXPECT().
			LogIn(gomock.Any(), *entity.NewUserCredentials("john", "strongpw")).
			Return(entity.AuthToken(""), entity.ErrNoSuchUser)

		grpcsvc := New(&dummyLogger{}, nil, uc, nil)

		_, err := grpcsvc.LogIn(ctx,
			&proto.LogInRequest{
				Credentials: &proto.UserCredentials{
					Login:    "john",
					Password: "strongpw",
				},
			},
		)

		assert.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("wrong password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockLogIner(ctrl)

		uc.EXPECT().
			LogIn(gomock.Any(), *entity.NewUserCredentials("john", "strongpw")).
			Return(entity.AuthToken(""), entity.ErrWrongPassword)

		grpcsvc := New(&dummyLogger{}, nil, uc, nil)

		_, err := grpcsvc.LogIn(ctx,
			&proto.LogInRequest{
				Credentials: &proto.UserCredentials{
					Login:    "john",
					Password: "strongpw",
				},
			},
		)

		assert.Equal(t, codes.Unauthenticated, status.Code(err))
	})
}

func TestStore(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull store", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockIStoreUC(ctrl)

		uc.EXPECT().Store(
			gomock.Any(),
			entity.NewStoreIn(
				*entity.NewUserCredentials("john", "strongpw"),
				"paycard",
				"tinkoff",
				`{"valid":true}`,
				[]byte("4437"),
			),
		)

		grpcsvc := New(&dummyLogger{}, nil, nil, uc)

		_, err := grpcsvc.Store(ctx,
			&proto.StoreRequest{
				Credentials: &proto.UserCredentials{
					Login:    "john",
					Password: "strongpw",
				},
				Type:    "paycard",
				Name:    "tinkoff",
				Meta:    `{"valid":true}`,
				Payload: []byte("4437"),
			},
		)

		assert.NoError(t, err)
	})

	t.Run("incorrect auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockIStoreUC(ctrl)

		uc.EXPECT().Store(
			gomock.Any(),
			gomock.Any(),
		).Return(entity.ErrAuthFailed)

		grpcsvc := New(&dummyLogger{}, nil, nil, uc)

		_, err := grpcsvc.Store(ctx,
			&proto.StoreRequest{
				Credentials: &proto.UserCredentials{},
			},
		)

		assert.Equal(t, codes.Unauthenticated, status.Code(err))
	})

	t.Run("other error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mock.NewMockIStoreUC(ctrl)

		uc.EXPECT().Store(
			gomock.Any(),
			gomock.Any(),
		).Return(errors.New("some other error"))

		grpcsvc := New(&dummyLogger{}, nil, nil, uc)

		_, err := grpcsvc.Store(ctx,
			&proto.StoreRequest{
				Credentials: &proto.UserCredentials{},
			},
		)

		assert.Equal(t, codes.Internal, status.Code(err))
	})
}
