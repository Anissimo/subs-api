package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"subs-api/internal/app"
	"subs-api/internal/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx, cfg, logger); err != nil {
		logger.Error("application stopped with error", slog.Any("error", err))
		os.Exit(1)
	}
}
