package log

import (
	"log/slog"
)

var log *slog.Logger

type Logger struct {
	*slog.Logger
}

func GetLogger() Logger {
	return Logger{log}
}

func MustSetup() *slog.Logger {
	panic("make me")
}
