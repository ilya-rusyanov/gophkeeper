package main

import (
	"context"
	"os/signal"
	"syscall"

	logging "github.com/ilya-rusyanov/gophkeeper/internal/logger"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/config"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/grpcserver"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/grpcservice"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/postgres"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/repository/user"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/login"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/register"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/store"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/usecase/token"
)

func main() {
	log := logging.MustNew("info")

	config := config.New()
	config.MustParse(log)

	log = logging.MustNew(config.LogLevel)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	db, err := postgres.New(
		ctx,
		log,
		config.DSN,
	)
	if err != nil {
		log.Fatal("database error: %s", err.Error())
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Error("failed to close DB: %s", err.Error())
		}
	}()

	tokenBuilder := token.NewBuilder(config.TokenSigningKey)

	userRepo := user.New(db)

	registerUC := register.New(
		config.UserPasswordSalt,
		userRepo,
		log,
		config.DefaultTokenLifetime,
		tokenBuilder,
	)

	logInUC := login.New(
		userRepo,
		config.UserPasswordSalt,
		tokenBuilder,
		config.DefaultTokenLifetime,
	)

	storeUC := store.New()

	grpcService := grpcservice.New(
		log,
		registerUC,
		logInUC,
		storeUC,
	)

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
