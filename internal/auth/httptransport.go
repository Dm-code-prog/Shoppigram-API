package auth

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the auth service.
func MakeHandler(bs Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	opts = append(opts, AuthServerBefore...)

	createOrUpdateTgUser := makeCreateOrUpdateTgUserEndpoint(bs)
	createOrUpdateTgUser = authMw(createOrUpdateTgUser)

	createOrUpdateTgUserHandler := kithttp.NewServer(
		createOrUpdateTgUser,
		decodeCreateOrUpdateTgUserRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Put("/telegram", createOrUpdateTgUserHandler.ServeHTTP)

	return r
}

func decodeCreateOrUpdateTgUserRequest(_ context.Context, _ *http.Request) (interface{}, error) {
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
	case errors.Is(err, ErrorBadRequest):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorBadRequest
	case errors.Is(err, ErrorUserNotFound):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorUserNotFound
	case errors.Is(err, ErrorInitDataIsMissing):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInitDataIsMissing
	case errors.Is(err, ErrorInitDataNotFound):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInitDataNotFound
	case errors.Is(err, ErrorInitDataIsInvalid):
		w.WriteHeader(http.StatusUnauthorized)
		err = ErrorInitDataIsInvalid
	case errors.Is(err, ErrorInitDataIsEmpty):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInitDataIsEmpty
	case errors.Is(err, ErrorWebAppNotFound):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorWebAppNotFound
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err = ErrorInternal
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
