package gokithelper

import (
	"context"
	"net/http"
)

// DecodeEmptyRequest is a helper function that decodes an empty request.
func DecodeEmptyRequest(_ context.Context, _ *http.Request) (any, error) {
	return nil, nil
}
