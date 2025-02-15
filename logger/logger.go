package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *slog.Logger
)

func Log() *slog.Logger {
	once.Do(func() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})

	return logger
}
