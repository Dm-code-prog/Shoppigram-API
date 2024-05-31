package httputils

import (
	"context"
	"net/http"
)

func DecodeEmptyRequest(c context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}
