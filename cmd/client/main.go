package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/config"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/controller"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/gophkeepergw"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/storage/usercred"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register"
	"github.com/ilya-rusyanov/gophkeeper/internal/logger"
)

func main() {
	log := logger.MustNew("info")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("configuration error: %q", err.Error())
	}

	log = logger.MustNew(cfg.LogLevel)

	userCredentialsStorage := usercred.New(
		log,
		"username.cfg",
		"gophkeeper",
	)

	gophkeeperGateway := gophkeepergw.New(
		cfg.Server,
	)

	registerUseCase := register.New(
		userCredentialsStorage,
		gophkeeperGateway,
	)

	ctrl := controller.New(
		controller.WithRegister(registerUseCase),
	)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	if err := ctrl.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
