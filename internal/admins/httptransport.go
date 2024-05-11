package admins

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
)

// MakeHandler returns a handler for the admin service.
func MakeHandler(bs *Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	getMarketplaces := makeGetMarketplaces(bs)
	getMarketplaces = authMw(getMarketplaces)

	getMarketplacesHandler := kithttp.NewServer(
		getMarketplaces,
		decodeCreateOrUpdateTgUserRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/", getMarketplacesHandler.ServeHTTP)

	return r
}

func decodeCreateOrUpdateTgUserRequest(c context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrorInvalidAdminID):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInvalidAdminID
	case errors.Is(err, ErrorAdminNotFound):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorAdminNotFound
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err = ErrorInternal
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
