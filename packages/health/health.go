package health

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Health ...
type Health struct {
	port     int
	endpoint string
	logger   *log.Logger
	srv      *http.Server
}

// New creates new Health listener
func New(port int, endpoint string) *Health {
	h := Health{
		port:     port,
		endpoint: endpoint,
		logger:   log.New(os.Stdout, "health", log.Llongfile|log.LUTC|log.Ldate|log.Ltime),
	}
	mux := http.NewServeMux()
	mux.HandleFunc(h.endpoint, NewHandler())
	h.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", h.port),
		Handler: mux,
	}
	return &h
}

// NewHandler returns new health handler function
func NewHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "OK")
		if err != nil {
			log.Println(err)
		}
	}
}

// ListenAndServe starts underlying server
func (h *Health) ListenAndServe() error {
	return h.srv.ListenAndServe()
}

// Shutdown delegates shutdown to underlying server
func (h *Health) Shutdown() error {
	return h.srv.Shutdown(context.Background())
}
