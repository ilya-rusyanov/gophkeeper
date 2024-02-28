package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/controller"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/gophkeepergw"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/storage/usercred"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register"
	"github.com/ilya-rusyanov/gophkeeper/internal/logger"
)

func main() {
	cfg := controller.ReadConfig()

	log := logger.MustNew(cfg.LogLevel)

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

	ctrl := controller.New()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	if err := ctrl.Run(
		ctx,
		controller.WithRegister(registerUseCase),
	); err != nil {
		log.Fatal(err)
	}
}
