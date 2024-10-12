package app

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

// decodeGetShopRequest decodes the request for the GetShop endpoint.
// The ID can be either a UUID or a short name. The request is malformed if the ID is missing.
func decodeGetShopRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		webAppID        uuid.UUID
		webAppShortName string
	)

	id := chi.URLParam(r, "id")
	if id == "" {
		return GetShopRequest{}, ErrorInvalidWebAppID
	}

	webAppID, err := uuid.Parse(id)
	if err != nil {
		webAppShortName = id
	}

	return GetShopRequest{
		WebAppID:        webAppID,
		WebAppShortName: webAppShortName,
	}, nil
}

func decodeInvalidateShopCacheRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return InvalidateShopCacheRequest{}, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return InvalidateShopCacheRequest{}, errors.Wrap(ErrorInvalidWebAppID, "uuid.Parse")
	}

	return InvalidateShopCacheRequest{
		WebAppID: webAppUUID,
	}, nil

}

func decodeCreateOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorInvalidWebAppID
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrorBadRequest
	}

	req.WebAppID = webAppUUID
	return req, nil
}

func decodeGetOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetOrderRequest
	orderId := chi.URLParam(r, "order_id")
	if orderId == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(orderId)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.OrderId = asUUID

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}
