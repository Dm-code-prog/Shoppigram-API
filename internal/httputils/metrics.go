package httputils

import (
	"github.com/shoppigram-com/marketplace-api/packages/cloudwatchcollector"
	"net/http"
	"regexp"
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

func (r *responseInterceptor) StatusCode() int {
	if r.statusCode == 0 {
		return http.StatusOK
	}

	return r.statusCode
}

func replaceUUIDs(input string) string {
	// Regular expression pattern for UUIDs (case-insensitive)
	re := regexp.MustCompile(`(?i)\b[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\b`)

	// Replace all UUIDs with "{id}"
	result := re.ReplaceAllString(input, "{id}")

	return result
}

// MakeObservabilityMiddleware is a middleware that collects metrics for each request.
// and publishes to AWS CloudWatch periodically.
func MakeObservabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ri := &responseInterceptor{ResponseWriter: w}
		next.ServeHTTP(ri, r)

		cloudwatchcollector.Increment("http_requests", cloudwatchcollector.Dimensions{
			"method": r.Method,
			"path":   replaceUUIDs(r.URL.Path),
			"status": strconv.Itoa(ri.StatusCode()),
		})
	})
}
