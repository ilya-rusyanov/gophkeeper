package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server is an HTTP server
type Server struct {
	server    *http.Server
	errorChan chan error
}

// New constructs HTTP server
func New(listenAddr string, handler http.Handler, opts ...Opt) *Server {
	res := &Server{
		server: &http.Server{
			Addr:    listenAddr,
			Handler: handler,
		},
		errorChan: make(chan error, 1),
	}

	// apply options
	for _, opt := range opts {
		opt(res.server)
	}

	go func() {
		if res.server.TLSConfig == nil {
			res.errorChan <- res.server.ListenAndServe()
		} else {
			res.errorChan <- res.server.ListenAndServeTLS("", "")
		}

		close(res.errorChan)
	}()

	return res
}

// Error supplies server errors
func (s *Server) Error() <-chan error {
	return s.errorChan
}

// Stop interrupts server
func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}
	return nil
}
