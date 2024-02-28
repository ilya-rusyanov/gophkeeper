package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/controller"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/gophkeepergw"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/storage/usercred"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/store"
	"github.com/ilya-rusyanov/gophkeeper/internal/logger"
)

func main() {
	cmdlineArgs := os.Args[1:]

	cfg := controller.ReadConfig(cmdlineArgs)

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

	storeUseCase := store.New()

	ctrl := controller.New(cmdlineArgs)

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
		controller.WithStore(storeUseCase),
	); err != nil {
		log.Fatal(err)
	}
}
