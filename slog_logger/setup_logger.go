package slog_logger

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string, level slog.Level) *slog.Logger {
	const op = "main.setupLogger"

	var handler slog.Handler
	switch env {
	case envLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		handler = NewColorHandler(handler)
	case envDev, envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	slogLog := slog.New(handler)
	slogLog.Info(op, slog.String("mode", fmt.Sprintf("сервис запущен в режиме '%s'", env)))
	return slogLog
}
