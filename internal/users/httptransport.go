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
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			xInitData := request.Header.Get("X-Init-Data")
			return PutInitDataToContext(ctx, xInitData)
		}),
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			webAppID := chi.URLParam(request, "web_app_id")
			return PutWebAppIDToContext(ctx, webAppID)
		}),
	}

	createOrUpdateTgUser := makeCreateOrUpdateTgUserEndpoint(bs)
	createOrUpdateTgUser = MakeAuthMiddleware(bs)(createOrUpdateTgUser)

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
