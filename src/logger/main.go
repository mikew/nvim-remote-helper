package logger

import (
	"log/slog"
	"os"

	"nvrh/src/logger/prettylog"
)

func PrepareLogger(isDebug bool) {
	logLevel := slog.LevelInfo

	if isDebug {
		logLevel = slog.LevelDebug
	}

	log := slog.New(prettylog.New(
		&slog.HandlerOptions{
			Level:     logLevel,
			AddSource: isDebug,
		},
		prettylog.WithDestinationWriter(os.Stderr),
		prettylog.WithColor(),
	))

	slog.SetDefault(log)
}
