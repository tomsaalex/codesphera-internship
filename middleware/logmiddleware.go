package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type ctxKeyLogger struct{}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Attach a logger with request-specific fields
		logger := slog.With(
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("request_id", uuid.NewString()),
		)

		// Add logger to context
		ctx := context.WithValue(r.Context(), ctxKeyLogger{}, logger)

		// Pass context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(ctxKeyLogger{}).(*slog.Logger)
	if !ok || logger == nil {
		return slog.Default()
	}
	return logger
}
