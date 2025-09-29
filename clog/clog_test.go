package clog

import (
	"io"
	"log/slog"
	"testing"
)

func TestWithContext(t *testing.T) {

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	ctx := WithContext(t.Context(), logger)

	l := ctx.Value(loggerKey)
	if l == nil {
		t.Fatal("unexpected nil logger on context")
	}

	if l != logger {
		t.Fatalf("mismatch (-want, +got): (%v, %v)", logger, l)
	}
}

// Tests that using WithContext with a nil logger fallsback to the package's DefaultLogger
func TestWithContextFallback(t *testing.T) {

	ctx := WithContext(t.Context(), nil)

	l := ctx.Value(loggerKey)
	if l == nil {
		t.Fatal("unexpected nil logger on context")
	}

	if l != DefaultLogger {
		t.Fatalf("mismatch (-want, +got): (%v, %v)", DefaultLogger, l)
	}
}

func TestFromContext(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	ctx := WithContext(t.Context(), logger)

	l := FromContext(ctx)

	if l != logger {
		t.Fatalf("mismatch (-want, +got): (%v, %v)", logger, l)
	}
}

func TestFromContextFallback(t *testing.T) {

	l := FromContext(t.Context())

	if l != DefaultLogger {
		t.Fatalf("mismatch (-want, +got): (%v, %v)", DefaultLogger, l)
	}
}
