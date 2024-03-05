package log

import (
	"log/slog"
	"os"
)

var log *slog.Logger

type Logger struct {
	*slog.Logger
}

func GetLogger() Logger {
	return Logger{log}
}

func MustSetup() *slog.Logger {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return l
}
