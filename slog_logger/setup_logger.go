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

// Цветное логирование
func setupLogger(env string) *slog.Logger {
	const op = "main.setupLogger"
	var slogLog *slog.Logger

	switch env {
	case envLocal:
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
		colorHandler := NewColorHandler(handler)
		slogLog = slog.New(colorHandler)
		slogLog.Info(op, slog.String("mode", fmt.Sprintf("сервис запущен в режиме '%s'", envLocal)))
	case envDev:
		slogLog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		slogLog.Info(op, slog.String("mode", fmt.Sprintf("сервис запущен в режиме '%s'", envDev)))
	case envProd:
		slogLog = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		slogLog.Info(op, slog.String("mode", fmt.Sprintf("сервис запущен в режиме '%s'", envProd)))
	}
	return slogLog
}
