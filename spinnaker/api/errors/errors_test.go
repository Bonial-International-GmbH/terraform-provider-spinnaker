package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponseError(t *testing.T) {
	err := errors.New("the-error")
	notFoundResp := &http.Response{StatusCode: http.StatusNotFound}

	require.EqualError(t, NewResponseError(nil, err), "the-error")
	require.EqualError(t, NewResponseError(notFoundResp, err), "the-error, Code: 404")
}

func TestIsNotFound(t *testing.T) {
	err := errors.New("the-error")
	respErr := NewResponseError(nil, err)
	notFoundErr := NewResponseError(&http.Response{StatusCode: http.StatusNotFound}, err)

	require.False(t, IsNotFound(err))
	require.False(t, IsNotFound(respErr))
	require.True(t, IsNotFound(notFoundErr))
}
