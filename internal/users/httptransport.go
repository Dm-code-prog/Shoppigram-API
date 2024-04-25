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

	telegramAuthUser := makeTelegramAuthUserEndpoint(bs)
	telegramAuthUser = makeTelegramRequestValidationMiddleware(bs)(telegramAuthUser)

	telegramAuthUserHandler := kithttp.NewServer(
		telegramAuthUser,
		decodeTelegramAuthUserRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Put("/telegram_auth", telegramAuthUserHandler.ServeHTTP)

	return r
}

func decodeTelegramAuthUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var usr User

	param := chi.URLParam(r, "user")
	if !json.Valid([]byte(param)) {
		return TelegramAuthUserRequest{}, ErrorBadRequest
	}

	err := json.Unmarshal([]byte(param), &usr)
	if err != nil {
		return TelegramAuthUserRequest{}, ErrorInternal
	}

	if usr.ExternalId == 0 {
		return TelegramAuthUserRequest{}, ErrorBadRequest
	}

	return TelegramAuthUserRequest{
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
