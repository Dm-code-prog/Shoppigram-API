package telegram_users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the users service.
func MakeHandler(bs *Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	createOrUpdateTelegramUser := makeCreateOrUpdateTelegramUserEndpoint(bs)
	createOrUpdateTelegramUser = makeValidateTelegramUserMiddleware(bs)(createOrUpdateTelegramUser)

	createOrUpdateTelegramUserHandler := kithttp.NewServer(
		createOrUpdateTelegramUser,
		decodeCreateOrUpdateTelegramUserRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Put("/telegram", createOrUpdateTelegramUserHandler.ServeHTTP)

	return r
}

// getUserFromContext gets User data from context
func getUserFromContext(ctx context.Context) (User, error) {
	val := ctx.Value("user")

	usr, ok := val.(User)
	if !ok {
		return User{}, ErrorInternal
	}
	return usr, nil
}

func decodeCreateOrUpdateTelegramUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	usr, err := getUserFromContext(c)
	if err != nil {
		return CreateOrUpdateTelegramUserRequest{}, errors.Wrap(err, "getUserFromContext")
	}

	return CreateOrUpdateTelegramUserRequest{
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
