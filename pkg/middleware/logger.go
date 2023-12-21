package middleware

import (
	"log/slog"
	"net/http"
)

func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var requestIDAttr slog.Attr
			requestID := r.Context().Value(RequestIDKey)
			if requestID != nil {
				requestIDAttr = slog.String("request-id", requestID.(string))
			}
			logger.Info("request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				requestIDAttr,
				slog.String("referer", r.Referer()),
			)
			next.ServeHTTP(w, r)
		})
	}
}
