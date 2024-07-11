package httputils

import (
	"net/http"

	"go.uber.org/zap"
)

type statusCodeResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *statusCodeResponseWriter) StatusCode() int {
	return w.statusCode
}

// MakeLoggingMiddleware creates a middleware that logs the request path, method, status code, and host.
func MakeLoggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sw := &statusCodeResponseWriter{w, http.StatusOK}
			next.ServeHTTP(sw, r)

			path := r.URL.Path
			method := r.Method
			statusCode := sw.StatusCode()
			logger.Info(
				"http request",
				zap.String("path", path),
				zap.String("method", method),
				zap.Int("status_code", statusCode),
				zap.String("host", r.Host),
			)
		})
	}
}
