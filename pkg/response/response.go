package response

import "fmt"

var MessageSuccess = "Success"

type Response struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ResponseError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	ERR        error
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

func WrapResponse(err error, code, status int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		ERR:        err,
		StatusCode: status,
	}
	return res
}

func Wrap400Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 0, 400, msg, args...)
}

func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 0, 500, msg, args...)
}

func NewResponse(code, status int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		StatusCode: status,
	}
	return res
}

func New400Response(msg string, args ...interface{}) error {
	return NewResponse(0, 400, msg, args...)
}

func New500Response(msg string, args ...interface{}) error {
	return NewResponse(0, 500, msg, args...)
}
