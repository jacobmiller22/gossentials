/*
 * clog is a useful package for storing and accessing a `log/slog` logger via a `context.Context`
 */
package clog

import (
	"context"
	"log/slog"
	"os"
)

// the context key type, used to avoid collisions
type key struct{}

var loggerKey key

// WithContext attaches the provided *slog.Logger to the provided context
// if the given logger is nil, the `DefaultLogger` will be attached
func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	if l == nil {
		l = DefaultLogger
	}

	return context.WithValue(ctx, loggerKey, l)
}

// FromContext retrieves a *slog.Logger from the given context
// If the retrieved logger isn't valid, the DefaultLogger will be retreived
func FromContext(ctx context.Context) *slog.Logger {
	l := ctx.Value(loggerKey)

	if l, ok := l.(*slog.Logger); ok && l != nil {
		return l
	}

	// Fallback
	return DefaultLogger
}

var DefaultLogger *slog.Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
