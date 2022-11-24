// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// Whoami implements whoami operation.
//
// GET /whoami
func (UnimplementedHandler) Whoami(ctx context.Context) (r User, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates ErrorResponseStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r ErrorResponseStatusCode) {
	return r
}
