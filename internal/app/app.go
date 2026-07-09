package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"subs-api/internal/config"
	apphttp "subs-api/internal/http"
)

func Run(ctx context.Context, cfg config.Config, logger *slog.Logger) error {
	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           apphttp.NewRouter(logger),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("http server started", slog.String("addr", cfg.HTTPAddr))
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}

		logger.Info("http server stopped")
		return nil
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
