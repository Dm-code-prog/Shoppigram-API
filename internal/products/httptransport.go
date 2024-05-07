package products

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs *Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndpoint(bs),
		decodeGetProductsRequest,
		encodeResponse,
		opts...,
	)

	invalidateProductsCacheHandler := kithttp.NewServer(
		makeInvalidateProductsCacheEndpoint(bs),
		decodeInvalidateProductsCacheRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/{web_app_id}", getProductsHandler.ServeHTTP)
	r.Put("/{web_app_id}/invalidate", invalidateProductsCacheHandler.ServeHTTP)

	return r
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return GetProductsRequest{}, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return GetProductsRequest{}, errors.Wrap(ErrorInvalidWebAppID, "uuid.Parse")
	}

	return GetProductsRequest{
		WebAppID: webAppUUID,
	}, nil
}

func decodeInvalidateProductsCacheRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return InvalidateProductsCacheRequest{}, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return InvalidateProductsCacheRequest{}, errors.Wrap(ErrorInvalidWebAppID, "uuid.Parse")
	}

	return InvalidateProductsCacheRequest{
		WebAppID: webAppUUID,
	}, nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrorNotFound):
		w.WriteHeader(http.StatusNotFound)
		err = ErrorNotFound
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
