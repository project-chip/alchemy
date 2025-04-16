package cmd

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
)

func configureLogging() {
	level := slog.LevelInfo
	switch strings.ToLower(commands.LogLevel) {
	case "error":
		level = slog.LevelError
	case "warn":
		level = slog.LevelWarn
	case "info":
		level = slog.LevelInfo
	case "debug":
		level = slog.LevelDebug
	}
	if commands.Verbose {
		level = slog.LevelDebug
	}

	var handler slog.Handler
	switch commands.Log {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	default:
		handler = tint.NewHandler(os.Stderr, &tint.Options{
			Level:      level,
			TimeFormat: time.StampMilli,
		})
	}
	if commands.ErrorExitCode {
		handler = &errorHandler{Handler: handler}
	}
	slog.SetDefault(slog.New(handler))
}
