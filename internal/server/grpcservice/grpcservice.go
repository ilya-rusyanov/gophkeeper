package grpcservice

import (
	"context"
	"fmt"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

type Logger interface {
	Debugf(string, ...any)
}

// RegisterUC is registration use case
type RegisterUC interface {
	Register(context.Context, entity.UserCredentials) error
}

// Service is a gophkeeper gRPC service
type Service struct {
	proto.UnimplementedGophkeeperServer
	log          Logger
	registration RegisterUC
}

// New constructs Service
func New(log Logger, reg RegisterUC) *Service {
	return &Service{
		log:          log,
		registration: reg,
	}
}

// Register is user registration endpoint
func (s *Service) Register(ctx context.Context, credentials *proto.UserCredentials) (*proto.Empty, error) {
	var res proto.Empty

	if err := s.registration.Register(
		ctx,
		*entity.NewUserCredentials(credentials.Login, credentials.Password),
	); err != nil {
		return &res, fmt.Errorf("failed to register: %w", err)
	}

	s.log.Debugf("incoming register request for %q completed successfully", credentials.Login)

	return &res, nil
}
