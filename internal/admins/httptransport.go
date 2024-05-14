package admins

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
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

	getMarketplaces := makeGetMarketplacesEndpoint(bs)
	getMarketplaces = authMw(getMarketplaces)

	createMarketplace := makeCreateMarketplaceEndpoint(bs)
	createMarketplace = authMw(createMarketplace)

	updateMarketplace := makeUpdateMarketplaceEndpoint(bs)
	updateMarketplace = authMw(updateMarketplace)

	getMarketplacesHandler := kithttp.NewServer(
		getMarketplaces,
		decodeGetMarketplacesRequest,
		encodeResponse,
		opts...,
	)

	createMarketplaceHandler := kithttp.NewServer(
		createMarketplace,
		decodeCreateMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	updateMarketplaceHandler := kithttp.NewServer(
		updateMarketplace,
		decodeUpdateMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	createProductHandler := kithttp.NewServer(
		authMw(makeCreateProductEndpoint(bs)),
		decodeCreateProductRequest,
		encodeResponse,
		opts...,
	)

	updateProductHandler := kithttp.NewServer(
		authMw(makeUpdateProductEndpoint(bs)),
		decodeUpdateProductRequest,
		encodeResponse,
		opts...,
	)

	deleteProductHandler := kithttp.NewServer(
		authMw(makeDeleteProductEndpoint(bs)),
		decodeDeleteProductRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/", getMarketplacesHandler.ServeHTTP)
	r.Post("/", createMarketplaceHandler.ServeHTTP)
	r.Put("/{web_app_id}", updateMarketplaceHandler.ServeHTTP)

	r.Post("/products/{web_app_id}", createProductHandler.ServeHTTP)
	r.Put("/products/{web_app_id}", updateProductHandler.ServeHTTP)
	r.Delete("/products/{web_app_id}", deleteProductHandler.ServeHTTP)

	return r
}

func decodeGetMarketplacesRequest(c context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

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

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

var badRequestErrors = []error{
	ErrorBadRequest,
	ErrorInvalidShortName,
	ErrorInvalidName,
	ErrorNotUniqueShortName,
	ErrorAdminNotFound,
	ErrorMaxMarketplacesExceeded,
	ErrorOpNotAllowed,
	ErrorMaxProductsExceeded,
	ErrorInvalidProductCurrency,
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	for _, d := range badRequestErrors {
		if errors.Is(err, d) {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"error": d.Error(),
			})
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": ErrorInternal.Error(),
	})
}
