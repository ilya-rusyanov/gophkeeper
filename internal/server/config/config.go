package config

import (
	"flag"
	"os"
)

// Config - app configuration
type Config struct {
	ListenAddr string
	DSN        string
}

// New constructs configuration
func New() *Config {
	var res Config

	flag.StringVar(&res.ListenAddr, "a", ":8080", "address and port to listen on")
	flag.StringVar(&res.DSN, "d", "", "data source name")

	return &res
}

// MustParse parses configuration or dies
func (c *Config) MustParse() {
	if val := os.Getenv("SERVER_ADDRESS"); len(val) > 0 {
		c.ListenAddr = val
	}

	if val := os.Getenv("DATABASE_DSN"); val != "" {
		c.DSN = val
	}
}
