package webhooks

import (
	"errors"
)

var (
	ErrorOrderDoesntExist = errors.New("order does not exist")
	ErrorBadRequest       = errors.New("bad request")
	ErrorCantHandle       = errors.New("can not handle the request")
	ErrorWrongResponce    = errors.New("can not create proper responce")
	ErrorDatabaseError    = errors.New("database error")
	ErrorWrongFormat      = errors.New("wrong data format")
)
