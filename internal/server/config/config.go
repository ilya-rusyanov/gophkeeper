package config

import (
	"flag"
	"os"
	"strconv"
	"time"
)

type Logger interface {
	Fatalf(string, ...any)
}

// Config - app configuration
type Config struct {
	ListenAddr           string
	LogLevel             string
	DSN                  string
	Secure               bool
	UserPasswordSalt     string
	DefaultTokenLifetime time.Duration
	TokenSigningKey      string
	tokenLifeSec         int
}

// New constructs configuration
func New() *Config {
	var res Config

	flag.StringVar(&res.ListenAddr, "a", ":8080", "address and port to listen on")
	flag.StringVar(&res.DSN, "d", "", "data source name")
	flag.StringVar(&res.LogLevel, "l", "info", "log level")
	flag.BoolVar(&res.Secure, "s", false, "enable HTTPS")
	flag.StringVar(&res.UserPasswordSalt, "user-password-salt", "6OpOisIp", "salt for hashing user passwords")
	flag.IntVar(&res.tokenLifeSec, "token-lifetime", 5*60, "default lifetime for user authentication token")
	flag.StringVar(&res.TokenSigningKey, "token-signing-key", "Twot3QuiOp", "signing key for user token")

	return &res
}

// MustParse parses configuration or dies
func (c *Config) MustParse(log Logger) {
	flag.Parse()

	if val := os.Getenv("GOPHKEEPER_SERVER_ADDRESS"); len(val) > 0 {
		c.ListenAddr = val
	}

	if val := os.Getenv("DATABASE_DSN"); val != "" {
		c.DSN = val
	}

	if val := os.Getenv("LOG_LEVEL"); len(val) > 0 {
		c.LogLevel = val
	}

	if val := os.Getenv("ENABLE_HTTPS"); len(val) > 0 {
		c.Secure = true
	}

	if val := os.Getenv("USER_PASSWORD_SALT"); len(val) > 0 {
		c.UserPasswordSalt = val
	}

	if val := os.Getenv("DEFAULT_TOKEN_LIFETIME"); len(val) > 0 {
		var err error
		c.tokenLifeSec, err = strconv.Atoi(val)
		if err != nil {
			log.Fatalf("failed to read default token lifetime from environment variable: %s", err.Error())
		}
	}
	c.DefaultTokenLifetime = time.Duration(c.tokenLifeSec) * time.Second

	if val := os.Getenv("TOKEN_SIGNING_KEY"); len(val) > 0 {
		c.TokenSigningKey = val
	}
}
