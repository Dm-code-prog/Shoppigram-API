package telegram_users

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the users service.
func MakeHandler(bs *Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	authUserHandler := kithttp.NewServer(
		makeAuthUserEndpoint(bs),
		decodeAuthUserRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Put("/auth", authUserHandler.ServeHTTP)

	return r
}

func decodeAuthUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var usr User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usr); err != nil {
		return nil, ErrorBadRequest
	}

	// TODO: Add request validation

	return AuthUserRequest{
		User: usr,
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
	case errors.Is(err, ErrorBadRequest):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorBadRequest
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err = ErrorInternal
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
