package admin

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
	"net/http"
)

var (
	ErrorAdminNotFound = errors.New("admin not found")
	ErrorInternal      = errors.New("internal server error")
	ErrorBadRequest    = errors.New("bad request")

	ErrorNotUniqueShortName      = errors.New("not unique short name")
	ErrorInvalidShortName        = errors.New("invalid short name")
	ErrorInvalidName             = errors.New("invalid name")
	ErrorMaxMarketplacesExceeded = errors.New("max marketplaces exceeded")

	ErrorMaxProductsExceeded = errors.New("max products exceeded")

	ErrorOpNotAllowed = errors.New("operation not allowed")

	ErrorInvalidImageExtension = errors.New("invalid image extension, only png, jpg, jpeg are allowed")

	ErrorShopSyncNotSupported = errors.New("shop sync not supported")

	badRequestErrors = []error{
		ErrorBadRequest,
		ErrorInvalidShortName,
		ErrorInvalidName,
		ErrorNotUniqueShortName,
		ErrorAdminNotFound,
		ErrorMaxMarketplacesExceeded,
		ErrorInvalidImageExtension,
		ErrorMaxProductsExceeded,
		ErrorShopSyncNotSupported,
	}
)

func encodeErrorHTTP(_ context.Context, err error, w http.ResponseWriter) {
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

	if errors.Is(err, ErrorOpNotAllowed) {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": ErrorOpNotAllowed.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": ErrorInternal.Error(),
	})
}
