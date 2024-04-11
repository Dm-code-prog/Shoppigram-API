package products

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs *Service) http.Handler {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndpoint(bs),
		decodeGetProductsRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Handle("/{web_app_id}/products", getProductsHandler)
	return r
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return GetProductsRequest{}, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return GetProductsRequest{}, ErrorInvalidWebAppID
	}

	return GetProductsRequest{
		WebAppID: webAppUUID,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrorNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, ErrorInvalidWebAppID):
		w.WriteHeader(http.StatusBadRequest)
	default:
		err = ErrorInternal
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
