package app

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
	"net/http"
)

var (
	ErrorProductsNotFound       = errors.New("products not found")
	ErrorInternal               = errors.New("internal server error")
	ErrorInvalidWebAppID        = errors.New("invalid web app id")
	ErrorInvalidProductQuantity = errors.New("the product quantity must be greater than zero")
	ErrorBadRequest             = errors.New("the request is malformed")
	ErrorGetOrderNotPermitted   = errors.New("No previleges to get order")
	ErrorInvalidOrderType       = errors.New("invalid order type")
	ErrorOrderIsNotSupported    = errors.New("orders are not supported in this app")

	badRequestErrors = []error{
		ErrorProductsNotFound,
		ErrorInvalidWebAppID,
		ErrorInvalidProductQuantity,
		ErrorBadRequest,
		ErrorGetOrderNotPermitted,
		ErrorBadRequest,
		ErrorInvalidOrderType,
		ErrorOrderIsNotSupported,
	}
)

func encodeErrorForHTTP(_ context.Context, err error, w http.ResponseWriter) {
	for _, e := range telegramusers.AuthenticationErrors {
		if errors.Is(err, e) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"error": e.Error(),
			})
			return
		}
	}

	for _, e := range badRequestErrors {
		if errors.Is(err, e) {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"error": e.Error(),
			})
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": ErrorInternal.Error(),
	})
}
