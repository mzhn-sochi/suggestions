package logger

import (
	"log/slog"
	"os"
)

func New(level slog.Level) *slog.Logger {
	h := &ContextHandler{slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})}

	l := slog.New(h)

	return l
}
