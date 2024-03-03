package grpcservice

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	grpcctx "github.com/ilya-rusyanov/gophkeeper/internal/server/grpcserver/context"
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

type Logger interface {
	Debugf(string, ...any)
}

// IRegisterUC is registration use case
type IRegisterUC interface {
	Register(context.Context, entity.UserCredentials) (entity.AuthToken, error)
}

// LogIner is log in use case
type LogIner interface {
	LogIn(context.Context, entity.UserCredentials) (entity.AuthToken, error)
}

// IStoreUC is store data use case
type IStoreUC interface {
	Store(context.Context, *entity.StoreIn) error
}

// Lister is listing use case
type Lister interface {
	List(ctx context.Context, username string) (entity.DataListing, error)
}

// Service is a gophkeeper gRPC service
type Service struct {
	proto.UnimplementedGophkeeperServer
	log            Logger
	registrationUC IRegisterUC
	loginUC        LogIner
	storeUC        IStoreUC
	listUC         Lister
}

// New constructs Service
func New(
	log Logger,
	reg IRegisterUC,
	login LogIner,
	store IStoreUC,
	list Lister,
) *Service {
	return &Service{
		log:            log,
		registrationUC: reg,
		loginUC:        login,
		storeUC:        store,
		listUC:         list,
	}
}

// Register is user registration endpoint
func (s *Service) Register(
	ctx context.Context, request *proto.RegisterRequest,
) (*proto.Empty, error) {
	var res proto.Empty

	token, err := s.registrationUC.Register(
		ctx,
		toUserCredentials(request.Credentials),
	)
	switch {
	case errors.Is(err, entity.ErrUserAlreadyExists):
		return nil, status.Error(
			codes.AlreadyExists,
			"user already exists")
	case err != nil:
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("failed to register: %s", err.Error()))
	}

	header := metadata.New(map[string]string{
		"token": string(token),
	})
	grpc.SendHeader(ctx, header)

	s.log.Debugf(
		"incoming register request for %q completed successfully",
		request.Credentials.Login)

	return &res, nil
}

// LogIn is log user in endpoint
func (s *Service) LogIn(
	ctx context.Context, request *proto.LogInRequest,
) (*proto.Empty, error) {
	var res proto.Empty

	token, err := s.loginUC.LogIn(
		ctx,
		toUserCredentials(request.Credentials),
	)
	switch {
	case errors.Is(err, entity.ErrNoSuchUser):
		return nil, status.Error(
			codes.NotFound,
			"user not found")
	case errors.Is(err, entity.ErrWrongPassword):
		return nil, status.Error(
			codes.Unauthenticated,
			"wrong password")
	case err != nil:
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("failed to log in: %s", err.Error()))
	}

	header := metadata.New(map[string]string{
		"token": string(token),
	})
	grpc.SendHeader(ctx, header)

	s.log.Debugf(
		"incoming log in request for %q completed successfully",
		request.Credentials.Login)

	return &res, nil
}

// Store performs storage of user's data
func (s *Service) Store(
	ctx context.Context, request *proto.StoreRequest,
) (*proto.Empty, error) {
	var res proto.Empty

	login := ctx.Value(grpcctx.ContextKeyLogin).(string)

	err := s.storeUC.Store(
		ctx,
		entity.NewStoreIn(
			login,
			request.Type,
			request.Name,
			request.Meta,
			request.Payload,
		),
	)
	switch {
	case errors.Is(err, entity.ErrAuthFailed):
		return nil, status.Error(
			codes.Unauthenticated,
			fmt.Sprintf("auth failed: %s", err.Error()))
	case errors.Is(err, entity.ErrRecordAlreadyExists):
		return nil, status.Error(
			codes.AlreadyExists,
			"record already exists")
	case err != nil:
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("internal error: %s", err.Error()))
	}

	return &res, nil
}

// List performs listing of user's data
func (s *Service) List(
	ctx context.Context, request *proto.ListRequest,
) (*proto.ListResponse, error) {
	var res proto.ListResponse

	login := ctx.Value(grpcctx.ContextKeyLogin).(string)

	l, err := s.listUC.List(
		ctx,
		login,
	)
	switch {
	case errors.Is(err, entity.ErrAuthFailed):
		return nil, status.Error(
			codes.Unauthenticated,
			fmt.Sprintf("auth failed: %s", err.Error()))
	case errors.Is(err, entity.ErrNoSuchUser):
		return nil, status.Error(
			codes.AlreadyExists,
			"user with given login does not exist")
	case err != nil:
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("internal error: %s", err.Error()))
	}

	for _, e := range l {
		res.Entries = append(res.Entries,
			&proto.ListResponse_Entry{
				Type: e.Type,
				Name: e.Name,
			})
	}

	return &res, nil
}

func toUserCredentials(c *proto.UserCredentials) entity.UserCredentials {
	var res entity.UserCredentials

	if c == nil {
		return res
	}

	res.Login = c.Login
	res.Password = c.Password

	return res
}
