package errors

import (
	"github.com/hackerchai/threatbook-ip-web/pkg/response"
	"github.com/pkg/errors"
)

var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	ErrNoPerm           = response.NewResponse(1001, 401, "no permission")
	ErrNotFound         = response.NewResponse(1002, 404, "not found")
	ErrMethodNotAllow   = response.NewResponse(1003, 405, "method not allowed")
	ErrTooManyRequests  = response.NewResponse(2000, 429, "too many requests")
	ErrInternalServer   = response.NewResponse(3000, 504, "internal server error")
	ErrInvalidParameter = response.NewResponse(1000, 500, "invalid parameter")
	ErrBadRequest       = response.New400Response("bad request")
)
