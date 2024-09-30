package httputils

import (
	"github.com/shoppigram-com/marketplace-api/packages/cloudwatchcollector"
	"net/http"
	"strconv"
)

type responseInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseInterceptor) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// MakeMetricsMiddleware is a middleware that collects metrics for each request.
// and publishes to AWS CloudWatch periodically.
func MakeMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ri := &responseInterceptor{ResponseWriter: w}
		next.ServeHTTP(ri, r)

		cloudwatchcollector.Increment("http_requests", cloudwatchcollector.Dimensions{
			"method": r.Method,
			"path":   r.URL.Path,
			"status": strconv.Itoa(ri.statusCode),
		})
	})
}
