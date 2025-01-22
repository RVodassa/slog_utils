package slog_logger

import (
	"context"
	"fmt"
	"log/slog"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Blue   = "\033[34m" // Добавлен синий цвет
)

type ColorHandler struct {
	handler slog.Handler
}

func NewColorHandler(handler slog.Handler) *ColorHandler {
	return &ColorHandler{handler: handler}
}

func (h *ColorHandler) Handle(ctx context.Context, record slog.Record) error {
	var color string

	switch record.Level {
	case slog.LevelDebug:
		color = Green
	case slog.LevelInfo:
		color = Blue
	case slog.LevelWarn:
		color = Yellow
	case slog.LevelError:
		color = Red
	default:
		color = Reset
	}

	// Печатаем цвет и передаём запись базовому обработчику
	fmt.Printf("%s", color)
	err := h.handler.Handle(ctx, record)
	fmt.Printf("%s", Reset) // Сбрасываем цвет

	return err
}

func (h *ColorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewColorHandler(h.handler.WithAttrs(attrs))
}

func (h *ColorHandler) WithGroup(name string) slog.Handler {
	return NewColorHandler(h.handler.WithGroup(name))
}
