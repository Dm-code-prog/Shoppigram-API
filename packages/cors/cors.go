package cors

import (
	"net/http"
	"regexp"
)

// MakeCORSMiddleware creates a new middleware that sets CORS headers based on the allowed origins and patterns.
func MakeCORSMiddleware(allowedOrigins []string, allowedPatterns []*regexp.Regexp) func(http.Handler) http.Handler {
	// Convert the slice of allowedOrigins to a map for efficient lookup
	allowedOriginsMap := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		allowedOriginsMap[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Set CORS headers if the origin is explicitly allowed or matches any regex pattern
			if _, ok := allowedOriginsMap[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				for _, pattern := range allowedPatterns {
					if pattern.MatchString(origin) {
						w.Header().Set("Access-Control-Allow-Origin", origin)
						break
					}
				}
			}

			// Set other CORS headers
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Init-Data")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Proceed with the next handler
			next.ServeHTTP(w, r)
		})
	}
}
