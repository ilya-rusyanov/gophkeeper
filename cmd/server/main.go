package main

import (
	"context"

	"github.com/go-chi/chi"

	"github.com/ilya-rusyanov/gophkeeper/internal/logger"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/config"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/httpserver"
	"github.com/ilya-rusyanov/gophkeeper/internal/shutdown"
)

func main() {
	config := config.New()
	config.MustParse()

	ctx := context.Background()

	logger := logger.MustNew(config.LogLevel)

	r := chi.NewRouter()

	httpServer := httpserver.New(config.ListenAddr, r)

	done := shutdown.Wait(ctx, logger, httpServer)

	<-done
}
