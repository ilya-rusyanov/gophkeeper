package main

import (
	"context"
	"os/signal"
	"syscall"

	log "github.com/ilya-rusyanov/gophkeeper/internal/logger"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/config"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/grpcserver"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/grpcservice"
)

func main() {
	config := config.New()
	config.MustParse()

	log := log.MustNew(config.LogLevel)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	grpcService := grpcservice.New(log)

	grpcServer, err := grpcserver.New(config.ListenAddr, grpcService, log)
	if err != nil {
		log.Fatalf("failed to create gRPC server: %s", err.Error())
	}

	errCh := grpcServer.Run()

	select {
	case <-ctx.Done():
		log.Info("stopping on signal")
		grpcServer.Stop()
	case <-errCh:
		for e := range errCh {
			log.Errorf("error running server: %q", e.Error())
		}
		log.Info("server stopped")
	}
}
