package marketplaces

import "github.com/pkg/errors"

var (
	ErrorProductsNotFound       = errors.New("products not found")
	ErrorInternal               = errors.New("internal server error")
	ErrorInvalidWebAppID        = errors.New("invalid web app id")
	ErrorInvalidProductQuantity = errors.New("the product quantity must be greater than zero")
	ErrorBadRequest             = errors.New("the request is malformed")
	ErrorGetOrderNotPremited    = errors.New("No previleges to get order")
)
