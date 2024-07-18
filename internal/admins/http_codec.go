package admins

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func decodeCreateMarketplaceRequest(c context.Context, r *http.Request) (interface{}, error) {
	var request CreateMarketplaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrorBadRequest
	}

	return request, nil
}

func decodeUpdateMarketplaceRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodeDeleteMarketplaceRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodeCreateProductRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodeUpdateProductRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodeDeleteProductRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodeCreateProductImageUploadURLRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodeCreateMarketplaceUploadLogoURLRequest(c context.Context, r *http.Request) (interface{}, error) {
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

func decodePublishMarketplaceBannerToChannelRequest(c context.Context, r *http.Request) (interface{}, error) {
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
