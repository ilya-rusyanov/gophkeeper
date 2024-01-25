package logger

import "go.uber.org/zap"

// Logger - main logger
type Logger = zap.SugaredLogger

// MustNew costructs logger or dies
func MustNew(level string) *Logger {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		panic(err)
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return zl.Sugar()
}
