package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (ms *MiddlewareStore) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)

		ms.Logger.Info("HTTP",
			zap.String("env", ms.ENV),
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.Int("status", rw.statusCode),
			zap.Duration("latency", time.Since(start)),
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the HTTP status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
