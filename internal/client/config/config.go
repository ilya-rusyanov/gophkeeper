package config

import (
	"fmt"
	"os"
)

// Config is app configuration values
type Config struct {
	Server   string
	LogLevel string
}

// New creates new app configuration
func New() (Config, error) {
	var (
		res Config
		ok  bool
	)

	res.Server, ok = os.LookupEnv("GOPHKEEPER_SERVER_ADDR")
	if !ok {
		return res, fmt.Errorf("GOPHKEEPER_SERVER_ADDR environment variable is not set")
	}

	res.LogLevel, ok = os.LookupEnv("LOG_LEVEL")
	if !ok {
		res.LogLevel = "info"
	}

	return res, nil
}
