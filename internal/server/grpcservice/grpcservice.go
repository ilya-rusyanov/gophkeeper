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
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

type Logger interface {
	Debugf(string, ...any)
}

// IRegisterUC is registration use case
type IRegisterUC interface {
	Register(context.Context, entity.UserCredentials) (entity.AuthToken, error)
}

// IStoreUC is store data use case
type IStoreUC interface {
	Store(context.Context, *entity.StoreIn) error
}

// Service is a gophkeeper gRPC service
type Service struct {
	proto.UnimplementedGophkeeperServer
	log          Logger
	registration IRegisterUC
	store        IStoreUC
}

// New constructs Service
func New(log Logger, reg IRegisterUC, store IStoreUC) *Service {
	return &Service{
		log:          log,
		registration: reg,
		store:        store,
	}
}

// Register is user registration endpoint
func (s *Service) Register(
	ctx context.Context, request *proto.RegisterRequest,
) (*proto.Empty, error) {
	var res proto.Empty

	token, err := s.registration.Register(
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

// Store performs storage of user's data
func (s *Service) Store(
	ctx context.Context, request *proto.StoreRequest,
) (*proto.Empty, error) {
	var res proto.Empty

	err := s.store.Store(
		ctx,
		entity.NewStoreIn(
			toUserCredentials(request.Credentials),
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
	case err != nil:
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("internal error: %s", err.Error()))
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
