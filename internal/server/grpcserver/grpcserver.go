package grpcserver

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// Logger is an interface to log
type Logger interface {
	Infof(string, ...any)
}

// Server is an gRPC server
type Server struct {
	listener   net.Listener
	server     *grpc.Server
	logger     Logger
	listenAddr string
}

// New constructs gRPC server
func New(listenAddr string, logger Logger) (*Server, error) {
	var err error

	res := Server{
		logger:     logger,
		listenAddr: listenAddr,
	}

	res.listener, err = net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to start listener: %w", err)
	}

	logger.Infof("hello")
	res.server = grpc.NewServer()

	// TODO: register grpc service

	return &res, nil
}

// Run starts gRPC server
func (s *Server) Run() <-chan error {
	errCh := make(chan error)

	go func() {
		s.logger.Infof("starting gRPC server on %s", s.listenAddr)
		if err := s.server.Serve(s.listener); err != nil {
			errCh <- fmt.Errorf("error serving gRPC server: %w", err)
		}

		defer close(errCh)
	}()

	return errCh
}

// Stop interrupts server
func (s *Server) Stop() {
	s.server.GracefulStop()
}
