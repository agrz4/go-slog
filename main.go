package main

import (
	"log/slog"
	"math/rand"
	"os"
	"runtime"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}).
		WithAttrs([]slog.Attr{slog.String("app_version", "v1.0.0")}))
	slog.SetDefault(logger)

	slog.Info(
		"hey Go!",
		slog.String("version", runtime.Version()),
		slog.Int("Random number", rand.Int()),
	)
	slog.Error("Testing error")
	slog.Warn("Testing Warn")
	slog.Debug("Debugging!")
}
