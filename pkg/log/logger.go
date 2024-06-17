package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type slogKey struct{}

type Slogger struct {
	*slog.Logger
}

func NewSlogger(handler slog.Handler) *Slogger {
	return &Slogger{slog.New(handler)}
}

func NewJsonHandler(w io.Writer, lovel slog.Level) slog.Handler {
	return slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: lovel,
	})
}

func FromContext(ctx context.Context) *Slogger {
	slogger, ok := ctx.Value(slogKey{}).(*Slogger)
	if !ok {
		return NewSlogger(NewJsonHandler(os.Stdout, slog.LevelDebug))
	}

	return slogger
}

func IntoContext(ctx context.Context, logger *Slogger) context.Context {
	return context.WithValue(ctx, slogKey{}, logger)
}

func Debug(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Debug(msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Info(msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Warn(msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(msg, args...)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(msg, args...)
	os.Exit(1)
}
