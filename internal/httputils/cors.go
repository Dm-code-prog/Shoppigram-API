package httputils

import (
	"net/http"
	"regexp"
)

// List of allowed origins
var allowedOrigins = map[string]struct{}{
	"https://web-app.shoppigram.com":             {},
	"https://admin.shoppigram.com":               {},
	"https://web-app.shoppigram.ru":              {},
	"https://dev-marketplace-web-app.vercel.app": {},
	"https://web-app.shoppigram.dev":             {},
	"http://localhost:3000":                      {},
	"http://localhost:5173":                      {},
}

// Slice of regex patterns for matching dynamic subdomains or origins
var allowedPatterns = []*regexp.Regexp{
	regexp.MustCompile(`^https://[\w-]+dmcodeprogs-projects\.vercel\.app$`),
	regexp.MustCompile(`^https://[\w-]+shoppigram\.vercel\.app$`),
}

// CORSMiddleware checks the request's origin and sets CORS headers accordingly.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Set CORS headers if the origin is explicitly allowed or matches any regex pattern
		if _, ok := allowedOrigins[origin]; ok {
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
