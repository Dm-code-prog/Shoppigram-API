package webhooks

import (
	"errors"
)

var (
	ErrorInternalServerError = errors.New("internal server error")
	ErrorOrderDoesntExist    = errors.New("order does not exist")
	ErrorBadRequest          = errors.New("bad request")
)
