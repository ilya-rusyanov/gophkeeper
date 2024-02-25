package grpcservice

import (
	"context"

	"github.com/ilya-rusyanov/gophkeeper/proto"
)

// Service is a gophkeeper gRPC service
type Service struct {
	proto.UnimplementedGophkeeperServer
}

// New constructs Service
func New() *Service {
	return &Service{}
}

func (s *Service) Register(ctx context.Context, credentials *proto.UserCredentials) (*proto.Empty, error) {
	var res proto.Empty

	// TODO

	return &res, nil
}
