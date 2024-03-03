package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/controller"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/gophkeepergw"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/storage/auth"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/list"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/login"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/register"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/show"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/store"
	"github.com/ilya-rusyanov/gophkeeper/internal/logger"
)

func main() {
	cmdlineArgs := os.Args[1:]

	cfg := controller.ReadConfig(cmdlineArgs)

	log := logger.MustNew(cfg.LogLevel)

	gophkeeperGateway := gophkeepergw.New(
		cfg.Server,
		log,
	)

	myAuthStorage := auth.New("auth.token")

	registerUseCase := register.New(
		gophkeeperGateway,
		myAuthStorage,
	)

	logInUseCase := login.New(
		gophkeeperGateway,
		myAuthStorage,
	)

	storeUseCase := store.New(
		myAuthStorage,
		gophkeeperGateway,
	)

	listUseCase := list.New(
		gophkeeperGateway,
		myAuthStorage,
	)

	showUseCase := show.New()

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
		controller.WithLogIn(logInUseCase),
		controller.WithStore(storeUseCase),
		controller.WithList(listUseCase),
		controller.WithOutput(os.Stdout),
		controller.WithShow(showUseCase),
	); err != nil {
		log.Fatal(err)
	}
}
