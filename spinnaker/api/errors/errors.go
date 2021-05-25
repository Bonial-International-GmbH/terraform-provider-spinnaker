package errors

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	gateapi "github.com/spinnaker/spin/gateapi"
)

var pipelineAlreadyExistsRegexp = regexp.MustCompile(`.*A pipeline with name .* already exists.*`)

// IsPipelineAlreadyExists returns true if the error indicates that a pipeline
// already exists.
func IsPipelineAlreadyExists(err error) bool {
	if err == nil {
		return false
	}

	return pipelineAlreadyExistsRegexp.MatchString(err.Error())
}

// IsNotFound returns true if err resembles an HTTP NotFound error.
func IsNotFound(err error) bool {
	return HasCode(http.StatusNotFound, err)
}

// HasCode returns true if err resembles an HTTP error with status code.
func HasCode(code int, err error) bool {
	var respErr *ResponseError

	if errors.As(err, &respErr) {
		return respErr.Code() == code
	}

	return false
}

// ResponseError wraps a (potentially nil) *http.Response and an error.
type ResponseError struct {
	resp *http.Response
	err  error
}

// NewResponseError creates a new *ResponseError.
func NewResponseError(resp *http.Response, err error) *ResponseError {
	return &ResponseError{
		resp: resp,
		err:  err,
	}
}

// Code returns the HTTP status code if the error includes an *http.Response.
// Otherwise returns 0.
func (e *ResponseError) Code() int {
	if e.resp != nil {
		return e.resp.StatusCode
	}

	return 0
}

// Error implements the error interface.
func (e *ResponseError) Error() string {
	if e.resp == nil {
		return e.err.Error()
	}

	code := e.Code()

	var gateErr gateapi.GenericSwaggerError

	if errors.As(e.err, &gateErr) {
		return fmt.Sprintf("%v, Code: %d, Body: %s", gateErr, code, string(gateErr.Body()))
	}

	return fmt.Sprintf("%v, Code: %d", e.err, code)
}
