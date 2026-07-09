package http

import (
	"log/slog"
	"net/http"

	"subs-api/internal/http/handlers"
	"subs-api/internal/http/middleware"
)

func NewRouter(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	handler := middleware.Recoverer(logger)(mux)
	handler = middleware.Logger(logger)(handler)

	return handler
}
