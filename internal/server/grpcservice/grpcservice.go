package grpcservice

import (
	"context"

	"github.com/ilya-rusyanov/gophkeeper/proto"
)

type Logger interface {
	Debugf(string, ...any)
}

// Service is a gophkeeper gRPC service
type Service struct {
	proto.UnimplementedGophkeeperServer
	log Logger
}

// New constructs Service
func New(log Logger) *Service {
	return &Service{
		log: log,
	}
}

func (s *Service) Register(ctx context.Context, credentials *proto.UserCredentials) (*proto.Empty, error) {
	var res proto.Empty

	// TODO

	s.log.Debugf("incoming register request for %q completed successfully", credentials.Login)

	return &res, nil
}
