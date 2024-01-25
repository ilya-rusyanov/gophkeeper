package main

import (
	"context"

	"github.com/ilya-rusyanov/gophkeeper/internal/logger"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/config"
	"github.com/ilya-rusyanov/gophkeeper/internal/shutdown"
)

func main() {
	config := config.New()
	config.MustParse()

	ctx := context.Background()

	logger := logger.MustNew(config.LogLevel)

	done := shutdown.Wait(ctx, logger, nil...)

	<-done
}
