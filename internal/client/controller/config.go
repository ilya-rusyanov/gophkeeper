package controller

// Config is app configuration values
type Config struct {
	// Server is remote server addr
	Server string `short:"r" name:"remote" env:"GOPHKEEPER_SERVER_ADDR"`
	// LogLevel is log level verbosity
	LogLevel string `short:"v" name:"verbosity" env:"LOG_LEVEL" enum:"debug,info,warn,error" default:"info"`
}
