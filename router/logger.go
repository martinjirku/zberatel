package router

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := &responseWriterWrapper{ResponseWriter: w}

			next.ServeHTTP(ww, r)

			duration := time.Since(start)

			reqID := middleware.GetReqID(r.Context())
			logger.Info("request handled",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.String("requestID", reqID),
				slog.Int("status", ww.status),
				slog.Duration("duration", duration),
			)
		})
	}
}

type responseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func (ww *responseWriterWrapper) WriteHeader(status int) {
	ww.status = status
	ww.ResponseWriter.WriteHeader(status)
}
