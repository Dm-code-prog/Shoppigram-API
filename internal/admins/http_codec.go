package admins

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func decodeCreateMarketplaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateMarketplaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	return request, nil
}

func decodeUpdateMarketplaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateMarketplaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	id := chi.URLParam(r, "web_app_id")
	if id == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.ID = asUUID

	return request, nil
}

func decodeDeleteMarketplaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteMarketplaceRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	id := chi.URLParam(r, "web_app_id")
	if id == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppId = asUUID

	return request, nil
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateProductRequest

	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppID = asUUID

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	return request, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppID = asUUID

	return request, nil
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppID = asUUID

	return request, nil
}

func decodeGetOrdersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// get all data from the query params
	var request GetOrdersRequest

	marketplaceID := r.URL.Query().Get("marketplace_id")
	if marketplaceID != "" {
		marketplaceUUID, err := uuid.Parse(marketplaceID)
		if err != nil {
			return nil, ErrorBadRequest
		}
		request.MarketplaceID = marketplaceUUID
	}

	state := r.URL.Query().Get("state")
	if state != "" {
		request.State = state
	}

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, ErrorBadRequest
		}

		request.Limit = limitInt
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return nil, ErrorBadRequest
		}

		request.Offset = offsetInt
	}

	return request, nil
}

func decodeCreateProductImageUploadURLRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateProductImageUploadURLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppID = asUUID

	return request, nil
}

func decodeCreateMarketplaceUploadLogoURLRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateMarketplaceLogoUploadURLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppID = asUUID

	return request, nil
}

func decodePublishMarketplaceBannerToChannelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PublishMarketplaceBannerToChannelRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.WebAppID = asUUID

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
